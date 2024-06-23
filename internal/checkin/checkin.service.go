package checkin

import (
	"context"

	"github.com/isd-sgcu/rpkm67-checkin/internal/model"
	userRepo "github.com/isd-sgcu/rpkm67-checkin/internal/user"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	proto.CheckInServiceServer
}

type serviceImpl struct {
	proto.UnimplementedCheckInServiceServer
	repo     Repository
	userRepo userRepo.Repository
	log      *zap.Logger
}

func NewService(repo Repository, log *zap.Logger) Service {
	return &serviceImpl{repo: repo, log: log}
}

func (s *serviceImpl) Create(_ context.Context, req *proto.CreateCheckInRequest) (*proto.CreateCheckInResponse, error) {
	var user model.User
	if err := s.userRepo.FindByEmail(req.Email, &user); err != nil {
		s.log.Named("Create").Error("FindByEmail: ", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}
	checkin := &model.Checkin{
		Email: req.Email,
		Event: req.Event,
	}

	err := s.repo.Create(checkin)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.CreateCheckInResponse{
		CheckIn: ModelToProto(checkin),
	}, nil
}

func (s *serviceImpl) FindByEmail(_ context.Context, req *proto.FindByEmailCheckInRequest) (*proto.FindByEmailCheckInResponse, error) {
	var checkins []*model.Checkin
	if err := s.repo.FindByEmail(req.Email, &checkins); err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.FindByEmailCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}

func ModelToProto(in *model.Checkin) *proto.CheckIn {
	return &proto.CheckIn{
		Email: in.Email,
		Event: in.Event,
	}
}

func ModelToProtoList(in []*model.Checkin) []*proto.CheckIn {
	var out []*proto.CheckIn
	for _, v := range in {
		out = append(out, ModelToProto(v))
	}
	return out
}
