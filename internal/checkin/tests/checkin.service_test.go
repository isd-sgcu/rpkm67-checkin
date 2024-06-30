package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-checkin/internal/checkin"
	mock_checkin "github.com/isd-sgcu/rpkm67-checkin/mocks/checkin"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/checkin/checkin/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CheckinServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	logger *zap.Logger
	checkinsProto []*proto.CheckIn
	checkinProto *proto.CheckIn
	createCheckInProtoRequest *proto.CreateCheckInRequest
	findByEmailCheckInRequest *proto.FindByEmailCheckInRequest
	findByUserIdCheckInRequest *proto.FindByUserIdCheckInRequest
}

func TestPinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
}

func (t *CheckinServiceTest) TestCreateSuccess() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)

	expected := &proto.CreateCheckInResponse {
		CheckIn: &proto.CheckIn {
			Id: "1", UserId: "1", Email: "1", Event: "1",
 		},
	}

}

func (t *CheckinServiceTest) TestCreateFailed() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)
}

func (t *CheckinServiceTest) TestFindByEmailSuccess() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)
}

func (t *CheckinServiceTest) TestFindByEmailFailed() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)
}

func (t *CheckinServiceTest) TestFindByUserIdSuccess() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)
}

func (t *CheckinServiceTest) TestFindByUserIdFailed() {
	repo := mock_checkin.NewMockRepository(t.controller)
	svc := checkin.NewService(repo, t.logger)
}
