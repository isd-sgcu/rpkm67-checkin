package test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-checkin/internal/checkin"
	mock_checkin "github.com/isd-sgcu/rpkm67-checkin/mocks/checkin"
	"github.com/isd-sgcu/rpkm67-gateway/apperror"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/isd-sgcu/rpkm67-model/model"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckinServiceTest struct {
	suite.Suite
	controller 				   *gomock.Controller
	logger 					   *zap.Logger
	checkinModel			   *model.CheckIn
	checkinsProto 			   []*proto.CheckIn
	checkinProto 			   *proto.CheckIn
	createCheckInProtoRequest  *proto.CreateCheckInRequest
	findByEmailCheckInRequest  *proto.FindByEmailCheckInRequest
	findByUserIdCheckInRequest *proto.FindByUserIdCheckInRequest
}

func TestPinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
	t.checkinModel = MockCheckInModel()
	t.checkinProto = MockCheckInProto()
	t.checkinsProto = MockCheckInsProto()
	t.createCheckInProtoRequest = &proto.CreateCheckInRequest{
		Email: t.checkinProto.Email,
		UserId: t.checkinProto.UserId,
		Event: t.checkinProto.Event,
	}
	t.findByEmailCheckInRequest = &proto.FindByEmailCheckInRequest{
		Email: t.checkinProto.Email,
	}
	t.findByUserIdCheckInRequest = &proto.FindByUserIdCheckInRequest{
		UserId: t.checkinProto.UserId,
	}
}

func (t *CheckinServiceTest) TestCreateSuccess() {
	repo := mock_checkin.NewMockRepository(t.controller)	
	svc := checkin.NewService(repo, t.logger)

	expectedResp := &proto.CreateCheckInResponse{
		CheckIn: t.checkinProto,
	}
	
	repo.EXPECT().Create(t.checkinModel).Return(nil)

	actual, err := svc.Create(context.Background(), t.createCheckInProtoRequest)

	t.Equal(actual, expectedResp)
	t.Nil(err)

}

func (t *CheckinServiceTest) TestCreateInternalError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.Internal, apperror.InternalServer.Error())
	repo.EXPECT().Create(t.checkinModel).Return(expectedErr)

	res, err := svc.Create(context.Background(), t.createCheckInProtoRequest)

	t.Nil(res)
	t.Equal(apperror.InternalServer, err)

}

func (t *CheckinServiceTest) TestFindByEmailSuccess() {
	
}
