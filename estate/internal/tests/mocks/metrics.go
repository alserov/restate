// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\metrics\metrics.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// ObserveRequest mocks base method.
func (m *MockMetrics) ObserveRequest(ctx context.Context, status int, dur time.Duration, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveRequest", ctx, status, dur, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObserveRequest indicates an expected call of ObserveRequest.
func (mr *MockMetricsMockRecorder) ObserveRequest(ctx, status, dur, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveRequest", reflect.TypeOf((*MockMetrics)(nil).ObserveRequest), ctx, status, dur, name)
}
