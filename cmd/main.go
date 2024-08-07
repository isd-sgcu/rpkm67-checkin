package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/isd-sgcu/rpkm67-checkin/config"
	"github.com/isd-sgcu/rpkm67-checkin/database"
	"github.com/isd-sgcu/rpkm67-checkin/internal/checkin"
	"github.com/isd-sgcu/rpkm67-checkin/logger"
	checkinProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger := logger.New(conf)

	// tp, err := tracer.New(conf)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to create tracer: %v", err))
	// }
	// defer func() {
	// 	if err := tp.Shutdown(context.Background()); err != nil {
	// 		panic(fmt.Sprintf("Failed to shutdown tracer: %v", err))
	// 	}
	// }()

	// tracer := tp.Tracer("rpkm67-checkin")

	db, err := database.InitDatabase(&conf.Db, conf.App.IsDevelopment())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// checkinRepo := checkin.NewRepository(db, tracer)
	// checkinSvc := checkin.NewService(checkinRepo, logger.Named("checkinSvc"), tracer)
	checkinRepo := checkin.NewRepository(db)
	checkinSvc := checkin.NewService(checkinRepo, logger.Named("checkinSvc"))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	checkinProto.RegisterCheckInServiceServer(grpcServer, checkinSvc)

	reflection.Register(grpcServer)
	go func() {
		logger.Sugar().Infof("RPKM67 Checkin starting at port %v", conf.App.Port)

		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatal("Failed to start RPKM67 Checkin service", zap.Error(err))
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, logger, map[string]operation{
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			sqlDB, err := db.DB()
			if err != nil {
				return nil
			}
			return sqlDB.Close()
		},
	})

	<-wait

	grpcServer.GracefulStop()
	logger.Info("Closing the listener")
	listener.Close()
	logger.Info("RPKM67 CheckIn service has been shutdown gracefully")
}

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, log *zap.Logger, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-s

		log.Named("graceful shutdown").Sugar().
			Infof("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Named("graceful shutdown").Sugar().
				Errorf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Named("graceful shutdown").Sugar().
					Infof("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Named("graceful shutdown").Sugar().
						Errorf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Named("graceful shutdown").Sugar().
					Infof("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}
