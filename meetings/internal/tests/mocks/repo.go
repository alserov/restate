// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\db\repository.go

// Package mock_db is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	models "github.com/alserov/restate/meetings/internal/service/models"
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

// ArrangeMeeting mocks base method.
func (m_2 *MockRepository) ArrangeMeeting(ctx context.Context, m models.Meeting) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "ArrangeMeeting", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArrangeMeeting indicates an expected call of ArrangeMeeting.
func (mr *MockRepositoryMockRecorder) ArrangeMeeting(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArrangeMeeting", reflect.TypeOf((*MockRepository)(nil).ArrangeMeeting), ctx, m)
}

// CancelMeeting mocks base method.
func (m *MockRepository) CancelMeeting(ctx context.Context, parameter models.CancelMeetingParameter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelMeeting", ctx, parameter)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelMeeting indicates an expected call of CancelMeeting.
func (mr *MockRepositoryMockRecorder) CancelMeeting(ctx, parameter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelMeeting", reflect.TypeOf((*MockRepository)(nil).CancelMeeting), ctx, parameter)
}

// GetMeetingTimestamps mocks base method.
func (m *MockRepository) GetMeetingTimestamps(ctx context.Context, estateID string) ([]time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeetingTimestamps", ctx, estateID)
	ret0, _ := ret[0].([]time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingTimestamps indicates an expected call of GetMeetingTimestamps.
func (mr *MockRepositoryMockRecorder) GetMeetingTimestamps(ctx, estateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingTimestamps", reflect.TypeOf((*MockRepository)(nil).GetMeetingTimestamps), ctx, estateID)
}

// GetMeetingsByEstateID mocks base method.
func (m *MockRepository) GetMeetingsByEstateID(ctx context.Context, estateID string) ([]models.Meeting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeetingsByEstateID", ctx, estateID)
	ret0, _ := ret[0].([]models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingsByEstateID indicates an expected call of GetMeetingsByEstateID.
func (mr *MockRepositoryMockRecorder) GetMeetingsByEstateID(ctx, estateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingsByEstateID", reflect.TypeOf((*MockRepository)(nil).GetMeetingsByEstateID), ctx, estateID)
}

// GetMeetingsByPhoneNumber mocks base method.
func (m *MockRepository) GetMeetingsByPhoneNumber(ctx context.Context, phoneNumber string) ([]models.Meeting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeetingsByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].([]models.Meeting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeetingsByPhoneNumber indicates an expected call of GetMeetingsByPhoneNumber.
func (mr *MockRepositoryMockRecorder) GetMeetingsByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeetingsByPhoneNumber", reflect.TypeOf((*MockRepository)(nil).GetMeetingsByPhoneNumber), ctx, phoneNumber)
}
