// Code generated by MockGen. DO NOT EDIT.
// Source: lib/atomic/atomic.go

// Package mock_atomic is a generated GoMock package.
package mock_atomic

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	atomic "esb-test/library/atomic"
)

// MockAtomicSessionProvider is a mock of AtomicSessionProvider interface.
type MockAtomicSessionProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAtomicSessionProviderMockRecorder
}

// MockAtomicSessionProviderMockRecorder is the mock recorder for MockAtomicSessionProvider.
type MockAtomicSessionProviderMockRecorder struct {
	mock *MockAtomicSessionProvider
}

// NewMockAtomicSessionProvider creates a new mock instance.
func NewMockAtomicSessionProvider(ctrl *gomock.Controller) *MockAtomicSessionProvider {
	mock := &MockAtomicSessionProvider{ctrl: ctrl}
	mock.recorder = &MockAtomicSessionProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtomicSessionProvider) EXPECT() *MockAtomicSessionProviderMockRecorder {
	return m.recorder
}

// BeginSession mocks base method.
func (m *MockAtomicSessionProvider) BeginSession(ctx context.Context) (*atomic.AtomicSessionContext, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginSession", ctx)
	ret0, _ := ret[0].(*atomic.AtomicSessionContext)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginSession indicates an expected call of BeginSession.
func (mr *MockAtomicSessionProviderMockRecorder) BeginSession(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginSession", reflect.TypeOf((*MockAtomicSessionProvider)(nil).BeginSession), ctx)
}

// MockAtomicSession is a mock of AtomicSession interface.
type MockAtomicSession struct {
	ctrl     *gomock.Controller
	recorder *MockAtomicSessionMockRecorder
}

// MockAtomicSessionMockRecorder is the mock recorder for MockAtomicSession.
type MockAtomicSessionMockRecorder struct {
	mock *MockAtomicSession
}

// NewMockAtomicSession creates a new mock instance.
func NewMockAtomicSession(ctrl *gomock.Controller) *MockAtomicSession {
	mock := &MockAtomicSession{ctrl: ctrl}
	mock.recorder = &MockAtomicSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAtomicSession) EXPECT() *MockAtomicSessionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockAtomicSession) Commit(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockAtomicSessionMockRecorder) Commit(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockAtomicSession)(nil).Commit), ctx)
}

// Rollback mocks base method.
func (m *MockAtomicSession) Rollback(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockAtomicSessionMockRecorder) Rollback(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockAtomicSession)(nil).Rollback), ctx)
}
