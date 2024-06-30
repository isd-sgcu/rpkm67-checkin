package checkin

import (
	"context"

	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/isd-sgcu/rpkm67-model/model"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	err := s.repo.Create(checkin)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		return nil, status.Error(codes.Internal, apperror.InternalServer.Error())
	}

	return &proto.CreateCheckInResponse{
		CheckIn: ModelToProto(checkin),
	}, nil
}

func (s *serviceImpl) FindByEmail(_ context.Context, req *proto.FindByEmailCheckInRequest) (*proto.FindByEmailCheckInResponse, error) {
	var checkins []*model.CheckIn
	if err := s.repo.FindByEmail(req.Email, &checkins); err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &proto.FindByEmailCheckInResponse{
		CheckIns: ModelToProtoList(checkins),
	}, nil
}

func (s *serviceImpl) FindByUserId(_ context.Context, req *proto.FindByUserIdCheckInRequest) (*proto.FindByUserIdCheckInResponse, error) {
	var checkins []*model.CheckIn
	if err := s.repo.FindByUserId(req.UserId, &checkins); err != nil {
		s.log.Named("FindByUserId").Error("FindByUserId: ", zap.Error(err))
		return nil, status.Error(codes.Internal, "internal error")
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
