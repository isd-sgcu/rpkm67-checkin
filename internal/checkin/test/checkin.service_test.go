package test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-checkin/config"
	"github.com/isd-sgcu/rpkm67-checkin/constant"
	"github.com/isd-sgcu/rpkm67-checkin/internal/checkin"
	mock_checkin "github.com/isd-sgcu/rpkm67-checkin/mocks/checkin"
	"github.com/isd-sgcu/rpkm67-checkin/tracer"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/isd-sgcu/rpkm67-model/model"
	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckinServiceTest struct {
	suite.Suite
	controller                 *gomock.Controller
	logger                     *zap.Logger
	tracer                     trace.Tracer
	checkinsModel              []*model.CheckIn
	checkinModel               *model.CheckIn
	checkinsProto              []*proto.CheckIn
	checkinProto               *proto.CheckIn
	createCheckInProtoRequest  *proto.CreateCheckInRequest
	findByEmailCheckInRequest  *proto.FindByEmailCheckInRequest
	findByUserIdCheckInRequest *proto.FindByUserIdCheckInRequest
}

func TestCheckinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	tracer, err := tracer.New(&config.Config{})
	if err != nil {
		t.T().Fatal(err)
	}
	t.tracer = tracer.Tracer("test")
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
	t.checkinsModel = MockCheckInsModel()
	t.checkinModel = t.checkinsModel[0]
	t.checkinsProto = checkin.ModelToProtoList(t.checkinsModel, false)
	t.checkinProto = t.checkinsProto[0]
	t.createCheckInProtoRequest = &proto.CreateCheckInRequest{
		Email:  t.checkinProto.Email,
		UserId: t.checkinProto.UserId,
		Event:  t.checkinProto.Event,
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
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedResp := &proto.CreateCheckInResponse{
		CheckIn: t.checkinProto,
	}

	repo.EXPECT().FindByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().Create(gomock.Any(), t.checkinModel).Return(nil)

	res, err := svc.Create(context.Background(), t.createCheckInProtoRequest)

	t.Equal(res, expectedResp)
	t.Nil(err)

}

func (t *CheckinServiceTest) TestCreateInternalError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.Internal, constant.InternalServerErrorMessage)

	repo.EXPECT().FindByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedErr)

	res, err := svc.Create(context.Background(), t.createCheckInProtoRequest)

	t.Nil(res)
	t.Equal(err, expectedErr)

}

func (t *CheckinServiceTest) TestCreateAlreadyCheckinError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.AlreadyExists, constant.AlreadyCheckinErrorMessage)

	repo.EXPECT().FindByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().Create(gomock.Any(), t.checkinModel).Return(expectedErr)

	res, err := svc.Create(context.Background(), t.createCheckInProtoRequest)

	t.Nil(res)
	t.Equal(expectedErr, err)
}

func (t *CheckinServiceTest) TestCreateInvalidArgumentError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.InvalidArgument, constant.InvalidDataErrorMessage)

	repo.EXPECT().FindByUserId(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedErr)

	res, err := svc.Create(context.Background(), t.createCheckInProtoRequest)

	t.Nil(res)
	t.Equal(expectedErr, err)
}

func (t *CheckinServiceTest) TestFindByEmailSuccess() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedResp := &proto.FindByEmailCheckInResponse{
		CheckIns: t.checkinsProto,
	}

	email := t.checkinModel.Email

	repo.EXPECT().FindByEmail(gomock.Any(), email, gomock.Any()).SetArg(2, t.checkinsModel).Return(nil)

	res, err := svc.FindByEmail(context.Background(), t.findByEmailCheckInRequest)

	t.Nil(err)
	t.Equal(expectedResp, res)
}

func (t *CheckinServiceTest) TestFindByEmailInternalError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	email := t.checkinModel.Email

	expectedErr := status.Error(codes.Internal, constant.InternalServerErrorMessage)
	repo.EXPECT().FindByEmail(gomock.Any(), email, gomock.Any()).SetArg(2, t.checkinsModel).Return(expectedErr)

	res, err := svc.FindByEmail(context.Background(), t.findByEmailCheckInRequest)

	t.Nil(res)
	t.Equal(err, expectedErr)
}

func (t *CheckinServiceTest) TestFindByEmailRequestCanceledError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.Canceled, constant.RequestCancelledErrorMessage)
	repo.EXPECT().FindByEmail(gomock.Any(), t.checkinModel.Email, gomock.Any()).SetArg(2, t.checkinsModel).Return(expectedErr)

	res, err := svc.FindByEmail(context.Background(), t.findByEmailCheckInRequest)

	t.Nil(res)
	t.Equal(err, expectedErr)
}

func (t *CheckinServiceTest) TestFindByUserIdSuccess() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedResp := &proto.FindByUserIdCheckInResponse{
		CheckIns: t.checkinsProto,
	}

	repo.EXPECT().FindByUserId(gomock.Any(), t.checkinModel.UserID, gomock.Any()).SetArg(2, t.checkinsModel).Return(nil)

	res, err := svc.FindByUserId(context.Background(), t.findByUserIdCheckInRequest)

	t.Nil(err)
	t.Equal(expectedResp, res)
}

func (t *CheckinServiceTest) TestFindByUserIdInternalError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.Internal, constant.InternalServerErrorMessage)
	repo.EXPECT().FindByUserId(gomock.Any(), t.checkinModel.UserID, gomock.Any()).Return(expectedErr)

	res, err := svc.FindByUserId(context.Background(), t.findByUserIdCheckInRequest)

	t.Nil(res)
	t.Equal(err, expectedErr)
}

func (t *CheckinServiceTest) TestFindByUserIdRequestCanceledError() {
	repo := mock_checkin.NewMockRepository(t.controller)
	// svc := checkin.NewService(repo, t.logger, t.tracer)
	svc := checkin.NewService(repo, t.logger)

	expectedErr := status.Error(codes.Canceled, constant.RequestCancelledErrorMessage)
	repo.EXPECT().FindByUserId(gomock.Any(), t.checkinModel.UserID, gomock.Any()).Return(expectedErr)

	res, err := svc.FindByUserId(context.Background(), t.findByUserIdCheckInRequest)

	t.Nil(res)
	t.Equal(err, expectedErr)
}
