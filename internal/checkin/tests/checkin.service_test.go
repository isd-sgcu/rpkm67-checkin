package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-checkin/internal/checkin"
	mock_checkin "github.com/isd-sgcu/rpkm67-checkin/mocks/checkin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CheckinServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	logger *zap.Logger
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
