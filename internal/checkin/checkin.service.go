package checkin

import (
	"context"
	"errors"

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
		return nil, status.Error(codes.Internal, "internal error")
	}
	for _, v := range checkin_userIds {
		if v.Event == req.Event && v.UserID == req.UserId {
			return nil, status.Error(codes.AlreadyExists, "already checked in")
		}
	}
	err = s.repo.Create(checkin)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		if errors.Is(err, gorm.ErrInvalidDB) {
			return nil, status.Error(codes.Internal, "database connection error")
		}
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, status.Error(codes.InvalidArgument, "invalid data")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateCheckInResponse{
		CheckIn: ModelToProto(checkin),
	}, nil
}

func (s *serviceImpl) FindByEmail(_ context.Context, req *proto.FindByEmailCheckInRequest) (*proto.FindByEmailCheckInResponse, error) {
	if req.Email == "" {
		s.log.Named("FindByUserEmail").Error("FindByUserEmail: invalid user ID")
		return nil, status.Error(codes.InvalidArgument, "email cannot be empty")
	}

	var checkins []*model.CheckIn
	err := s.repo.FindByEmail(req.Email, &checkins)
	if err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "request canceled by the client")
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return nil, status.Error(codes.Internal, "database connection error")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timed out")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	if len(checkins) == 0 {
		s.log.Named("FindByUserEmail").Error("email not found")
		return nil, status.Error(codes.NotFound, "email not found")
	}

	return &proto.FindByEmailCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}

func (s *serviceImpl) FindByUserId(_ context.Context, req *proto.FindByUserIdCheckInRequest) (*proto.FindByUserIdCheckInResponse, error) {
	if req.UserId == "" {
		s.log.Named("FindByUserId").Error("FindByUserId: invalid user ID")
		return nil, status.Error(codes.InvalidArgument, "user ID cannot be empty")
	}

	var checkins []*model.CheckIn
	err := s.repo.FindByUserId(req.UserId, &checkins)
	if err != nil {
		s.log.Named("FindByUserId").Error("FindByUserId: ", zap.Error(err))
		if errors.Is(err, context.Canceled) {
			return nil, status.Error(codes.Canceled, "request canceled by the client")
		}
		if errors.Is(err, gorm.ErrInvalidDB) {
			return nil, status.Error(codes.Internal, "database connection error")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, status.Error(codes.DeadlineExceeded, "request timed out")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	if len(checkins) == 0 {
		s.log.Named("FindByUserId").Error("user not found")
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &proto.FindByUserIdCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}

func ModelToProto(in *model.CheckIn) *proto.CheckIn {
	return &proto.CheckIn{
		Id:     in.ID.String(),
		Email:  in.Email,
		Event:  in.Event,
		UserId: in.UserID,
	}
}

func ModelToProtoList(in []*model.CheckIn) []*proto.CheckIn {
	var out []*proto.CheckIn
	for _, v := range in {
		out = append(out, ModelToProto(v))
	}
	return out
}
