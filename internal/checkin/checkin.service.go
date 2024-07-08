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
}

func NewService(repo Repository, log *zap.Logger) Service {
	return &serviceImpl{repo: repo, log: log}
}

func (s *serviceImpl) Create(_ context.Context, req *proto.CreateCheckInRequest) (*proto.CreateCheckInResponse, error) {
	checkin := &model.CheckIn{
		Email:  req.Email,
		Event:  req.Event,
		UserID: req.UserId,
	}

	var checkin_userIds []*model.CheckIn
	err := s.repo.FindByUserId(req.UserId, &checkin_userIds)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}
	for _, v := range checkin_userIds {
		if v.Event == req.Event && v.UserID == req.UserId {
			s.log.Named("Create").Warn("Create: User already checkin this event")

			return &proto.CreateCheckInResponse{
				CheckIn: ModelToProto(checkin),
			}, nil
		}
	}
	err = s.repo.Create(checkin)
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

	return &proto.CreateCheckInResponse{
		CheckIn: ModelToProto(checkin),
	}, nil
}

func (s *serviceImpl) FindByEmail(_ context.Context, req *proto.FindByEmailCheckInRequest) (*proto.FindByEmailCheckInResponse, error) {
	if req.Email == "" {
		s.log.Named("FindByUserEmail").Error("FindByUserEmail: invalid user Email")
		return nil, status.Error(codes.InvalidArgument, constant.ArgumentEmptyErrorMessage)
	}

	var checkins []*model.CheckIn
	err := s.repo.FindByEmail(req.Email, &checkins)
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

	return &proto.FindByEmailCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}

func (s *serviceImpl) FindByUserId(_ context.Context, req *proto.FindByUserIdCheckInRequest) (*proto.FindByUserIdCheckInResponse, error) {
	if req.UserId == "" {
		s.log.Named("FindByUserId").Error("FindByUserId: invalid user ID")
		return nil, status.Error(codes.InvalidArgument, constant.ArgumentEmptyErrorMessage)
	}

	var checkins []*model.CheckIn
	err := s.repo.FindByUserId(req.UserId, &checkins)
	if err != nil {
		s.log.Named("FindByUserId").Error("FindByUserId: ", zap.Error(err))

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
