// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/checkin/checkin.repository.go

// Package mock_checkin is a generated GoMock package.
package mock_checkin

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/isd-sgcu/rpkm67-model/model"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, checkIn *model.CheckIn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, checkIn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, checkIn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, checkIn)
}

// FindByEmail mocks base method.
func (m *MockRepository) FindByEmail(ctx context.Context, email string, checkIns *[]*model.CheckIn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email, checkIns)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockRepositoryMockRecorder) FindByEmail(ctx, email, checkIns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockRepository)(nil).FindByEmail), ctx, email, checkIns)
}

// FindByUserId mocks base method.
func (m *MockRepository) FindByUserId(ctx context.Context, userId string, checkIns *[]*model.CheckIn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserId", ctx, userId, checkIns)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindByUserId indicates an expected call of FindByUserId.
func (mr *MockRepositoryMockRecorder) FindByUserId(ctx, userId, checkIns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserId", reflect.TypeOf((*MockRepository)(nil).FindByUserId), ctx, userId, checkIns)
}
