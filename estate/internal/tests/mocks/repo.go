// Code generated by MockGen. DO NOT EDIT.
// Source: .\estate\internal\db\repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/alserov/restate/estate/internal/service/models"
	gomock "github.com/golang/mock/gomock"
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

// CreateEstate mocks base method.
func (m *MockRepository) CreateEstate(ctx context.Context, estate models.Estate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEstate", ctx, estate)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEstate indicates an expected call of CreateEstate.
func (mr *MockRepositoryMockRecorder) CreateEstate(ctx, estate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEstate", reflect.TypeOf((*MockRepository)(nil).CreateEstate), ctx, estate)
}

// DeleteEstate mocks base method.
func (m *MockRepository) DeleteEstate(ctx context.Context, estateID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEstate", ctx, estateID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEstate indicates an expected call of DeleteEstate.
func (mr *MockRepositoryMockRecorder) DeleteEstate(ctx, estateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEstate", reflect.TypeOf((*MockRepository)(nil).DeleteEstate), ctx, estateID)
}

// GetEstateInfo mocks base method.
func (m *MockRepository) GetEstateInfo(ctx context.Context, estateID string) (models.Estate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateInfo", ctx, estateID)
	ret0, _ := ret[0].(models.Estate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateInfo indicates an expected call of GetEstateInfo.
func (mr *MockRepositoryMockRecorder) GetEstateInfo(ctx, estateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateInfo", reflect.TypeOf((*MockRepository)(nil).GetEstateInfo), ctx, estateID)
}

// GetEstateList mocks base method.
func (m *MockRepository) GetEstateList(ctx context.Context, param models.GetEstateListParameters) ([]models.EstateMainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateList", ctx, param)
	ret0, _ := ret[0].([]models.EstateMainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateList indicates an expected call of GetEstateList.
func (mr *MockRepositoryMockRecorder) GetEstateList(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateList", reflect.TypeOf((*MockRepository)(nil).GetEstateList), ctx, param)
}