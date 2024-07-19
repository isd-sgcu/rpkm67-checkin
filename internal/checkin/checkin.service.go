package checkin

import (
	"context"
	"errors"

	"github.com/isd-sgcu/rpkm67-checkin/constant"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/isd-sgcu/rpkm67-model/model"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Service interface {
	proto.CheckInServiceServer
}

type serviceImpl struct {
	proto.UnimplementedCheckInServiceServer
	repo Repository
	log  *zap.Logger
	// tracer trace.Tracer
}

func NewService(repo Repository, log *zap.Logger) Service {
	return &serviceImpl{repo: repo, log: log}
}

func (s *serviceImpl) Create(ctx context.Context, req *proto.CreateCheckInRequest) (*proto.CreateCheckInResponse, error) {
	// span := trace.SpanFromContext(ctx)
	// span.SetAttributes(
	// 	attribute.String("req.user.id", req.UserId),
	// 	attribute.String("req.user.email", req.Email),
	// 	attribute.String("req.event", req.Event),
	// )

	checkin := &model.CheckIn{
		Email:  req.Email,
		Event:  req.Event,
		UserID: req.UserId,
	}

	var checkin_userIds []*model.CheckIn
	err := s.repo.FindByUserId(ctx, req.UserId, &checkin_userIds)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}
	for _, v := range checkin_userIds {
		if v.Event == req.Event && v.UserID == req.UserId {
			s.log.Named("Create").Error("Create: User already checkin this event")

			return nil, status.Error(codes.AlreadyExists, constant.AlreadyCheckinErrorMessage)
		}
	}
	// span.AddEvent("Verify user checkin not duplicate")

	err = s.repo.Create(ctx, checkin)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		if status.Code(err) == codes.AlreadyExists {
			return nil, status.Error(codes.AlreadyExists, constant.AlreadyCheckinErrorMessage)
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return nil, status.Error(codes.Internal, constant.DatabaseConnectionErrorMessage)
		}
		if status.Code(err) == codes.InvalidArgument {
			return nil, status.Error(codes.InvalidArgument, constant.InvalidDataErrorMessage)
		}
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}
	// span.AddEvent("Checkin created")

	return &proto.CreateCheckInResponse{
		CheckIn: ModelToProto(checkin),
	}, nil
}

func (s *serviceImpl) FindByEmail(ctx context.Context, req *proto.FindByEmailCheckInRequest) (*proto.FindByEmailCheckInResponse, error) {
	// span := trace.SpanFromContext(ctx)
	// defer span.End()

	// span.SetAttributes(
	// 	attribute.String("req.user.email", req.Email),
	// )

	if req.Email == "" {
		s.log.Named("FindByUserEmail").Error("FindByUserEmail: invalid user Email")
		return nil, status.Error(codes.InvalidArgument, constant.ArgumentEmptyErrorMessage)
	}

	var checkins []*model.CheckIn
	err := s.repo.FindByEmail(ctx, req.Email, &checkins)
	if err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		if status.Code(err) == codes.Canceled {
			return nil, status.Error(codes.Canceled, constant.RequestCancelledErrorMessage)
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return nil, status.Error(codes.Internal, constant.DatabaseConnectionErrorMessage)
		}
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}
	// span.AddEvent("Checkin found")

	return &proto.FindByEmailCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}

func (s *serviceImpl) FindByUserId(ctx context.Context, req *proto.FindByUserIdCheckInRequest) (*proto.FindByUserIdCheckInResponse, error) {
	// ctx, span := s.tracer.Start(ctx, "service.checkin.FindByUserId")
	// span.SetAttributes(
	// 	attribute.String("req.user.id", req.UserId),
	// )
	// defer span.End()

	if req.UserId == "" {
		s.log.Named("FindByUserId").Error("FindByUserId: invalid user ID")
		// span.RecordError(status.Error(codes.InvalidArgument, constant.ArgumentEmptyErrorMessage))
		return nil, status.Error(codes.InvalidArgument, constant.ArgumentEmptyErrorMessage)
	}

	var checkins []*model.CheckIn
	err := s.repo.FindByUserId(ctx, req.UserId, &checkins)
	if err != nil {
		s.log.Named("FindByUserId").Error("FindByUserId: ", zap.Error(err))
		// span.RecordError(status.Error(codes.InvalidArgument, constant.ArgumentEmptyErrorMessage))

		if status.Code(err) == codes.Canceled {
			return nil, status.Error(codes.Canceled, constant.RequestCancelledErrorMessage)
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return nil, status.Error(codes.Internal, constant.DatabaseConnectionErrorMessage)
		}
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}

	return &proto.FindByUserIdCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}
