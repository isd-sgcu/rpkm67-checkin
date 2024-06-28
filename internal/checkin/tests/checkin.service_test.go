package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isd-sgcu/rpkm67-model/model"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CheckinServiceTest struct {
	suite.Suite
	controller *gomock.Controller
	Checkin *model.CheckIn
	logger *zap.Logger
	// More to be implemented
}

func TestPinService(t *testing.T) {
	suite.Run(t, new(CheckinServiceTest))
}

func (t *CheckinServiceTest) SetupTest() {
	t.controller = gomock.NewController(t.T())
	t.logger = zap.NewNop()
}

func (t *CheckinServiceTest) TestCreateSuccess() {

}

func (t *CheckinServiceTest) TestCreateFailed() {

}

func (t *CheckinServiceTest) TestFindByEmailSuccess() {

}

func (t *CheckinServiceTest) TestFindByEmailFailed() {
	
}

func (t *CheckinServiceTest) TestFindByUserIdSuccess() {
	
}

func (t *CheckinServiceTest) TestFindByUserIdFailed() {
	
}
