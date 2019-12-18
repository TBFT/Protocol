// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger-labs/minbft/core/internal/timer (interfaces: Timer,Provider)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	timer "github.com/hyperledger-labs/minbft/core/internal/timer"
	reflect "reflect"
	time "time"
)

// MockTimer is a mock of Timer interface
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *MockTimerMockRecorder
}

// MockTimerMockRecorder is the mock recorder for MockTimer
type MockTimerMockRecorder struct {
	mock *MockTimer
}

// NewMockTimer creates a new mock instance
func NewMockTimer(ctrl *gomock.Controller) *MockTimer {
	mock := &MockTimer{ctrl: ctrl}
	mock.recorder = &MockTimerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTimer) EXPECT() *MockTimerMockRecorder {
	return m.recorder
}

// Expired mocks base method
func (m *MockTimer) Expired() <-chan time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expired")
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

// Expired indicates an expected call of Expired
func (mr *MockTimerMockRecorder) Expired() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expired", reflect.TypeOf((*MockTimer)(nil).Expired))
}

// Reset mocks base method
func (m *MockTimer) Reset(arg0 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", arg0)
}

// Reset indicates an expected call of Reset
func (mr *MockTimerMockRecorder) Reset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockTimer)(nil).Reset), arg0)
}

// Stop mocks base method
func (m *MockTimer) Stop() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockTimerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTimer)(nil).Stop))
}

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AfterFunc mocks base method
func (m *MockProvider) AfterFunc(arg0 time.Duration, arg1 func()) timer.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterFunc", arg0, arg1)
	ret0, _ := ret[0].(timer.Timer)
	return ret0
}

// AfterFunc indicates an expected call of AfterFunc
func (mr *MockProviderMockRecorder) AfterFunc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterFunc", reflect.TypeOf((*MockProvider)(nil).AfterFunc), arg0, arg1)
}

// NewTimer mocks base method
func (m *MockProvider) NewTimer(arg0 time.Duration) timer.Timer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTimer", arg0)
	ret0, _ := ret[0].(timer.Timer)
	return ret0
}

// NewTimer indicates an expected call of NewTimer
func (mr *MockProviderMockRecorder) NewTimer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTimer", reflect.TypeOf((*MockProvider)(nil).NewTimer), arg0)
}
