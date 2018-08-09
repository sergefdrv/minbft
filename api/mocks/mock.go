// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger-labs/minbft/api (interfaces: Configer,Authenticator,RequestConsumer)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/hyperledger-labs/minbft/api"
	reflect "reflect"
	time "time"
)

// MockConfiger is a mock of Configer interface
type MockConfiger struct {
	ctrl     *gomock.Controller
	recorder *MockConfigerMockRecorder
}

// MockConfigerMockRecorder is the mock recorder for MockConfiger
type MockConfigerMockRecorder struct {
	mock *MockConfiger
}

// NewMockConfiger creates a new mock instance
func NewMockConfiger(ctrl *gomock.Controller) *MockConfiger {
	mock := &MockConfiger{ctrl: ctrl}
	mock.recorder = &MockConfigerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfiger) EXPECT() *MockConfigerMockRecorder {
	return m.recorder
}

// CheckpointPeriod mocks base method
func (m *MockConfiger) CheckpointPeriod() uint32 {
	ret := m.ctrl.Call(m, "CheckpointPeriod")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// CheckpointPeriod indicates an expected call of CheckpointPeriod
func (mr *MockConfigerMockRecorder) CheckpointPeriod() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckpointPeriod", reflect.TypeOf((*MockConfiger)(nil).CheckpointPeriod))
}

// F mocks base method
func (m *MockConfiger) F() uint32 {
	ret := m.ctrl.Call(m, "F")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// F indicates an expected call of F
func (mr *MockConfigerMockRecorder) F() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "F", reflect.TypeOf((*MockConfiger)(nil).F))
}

// Logsize mocks base method
func (m *MockConfiger) Logsize() uint32 {
	ret := m.ctrl.Call(m, "Logsize")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Logsize indicates an expected call of Logsize
func (mr *MockConfigerMockRecorder) Logsize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logsize", reflect.TypeOf((*MockConfiger)(nil).Logsize))
}

// N mocks base method
func (m *MockConfiger) N() uint32 {
	ret := m.ctrl.Call(m, "N")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// N indicates an expected call of N
func (mr *MockConfigerMockRecorder) N() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "N", reflect.TypeOf((*MockConfiger)(nil).N))
}

// TimeoutRequest mocks base method
func (m *MockConfiger) TimeoutRequest() time.Duration {
	ret := m.ctrl.Call(m, "TimeoutRequest")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeoutRequest indicates an expected call of TimeoutRequest
func (mr *MockConfigerMockRecorder) TimeoutRequest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutRequest", reflect.TypeOf((*MockConfiger)(nil).TimeoutRequest))
}

// TimeoutViewChange mocks base method
func (m *MockConfiger) TimeoutViewChange() time.Duration {
	ret := m.ctrl.Call(m, "TimeoutViewChange")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeoutViewChange indicates an expected call of TimeoutViewChange
func (mr *MockConfigerMockRecorder) TimeoutViewChange() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutViewChange", reflect.TypeOf((*MockConfiger)(nil).TimeoutViewChange))
}

// MockAuthenticator is a mock of Authenticator interface
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// GenerateMessageAuthenTag mocks base method
func (m *MockAuthenticator) GenerateMessageAuthenTag(arg0 api.AuthenticationRole, arg1 []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "GenerateMessageAuthenTag", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateMessageAuthenTag indicates an expected call of GenerateMessageAuthenTag
func (mr *MockAuthenticatorMockRecorder) GenerateMessageAuthenTag(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateMessageAuthenTag", reflect.TypeOf((*MockAuthenticator)(nil).GenerateMessageAuthenTag), arg0, arg1)
}

// VerifyMessageAuthenTag mocks base method
func (m *MockAuthenticator) VerifyMessageAuthenTag(arg0 api.AuthenticationRole, arg1 uint32, arg2, arg3 []byte) error {
	ret := m.ctrl.Call(m, "VerifyMessageAuthenTag", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyMessageAuthenTag indicates an expected call of VerifyMessageAuthenTag
func (mr *MockAuthenticatorMockRecorder) VerifyMessageAuthenTag(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMessageAuthenTag", reflect.TypeOf((*MockAuthenticator)(nil).VerifyMessageAuthenTag), arg0, arg1, arg2, arg3)
}

// MockRequestConsumer is a mock of RequestConsumer interface
type MockRequestConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockRequestConsumerMockRecorder
}

// MockRequestConsumerMockRecorder is the mock recorder for MockRequestConsumer
type MockRequestConsumerMockRecorder struct {
	mock *MockRequestConsumer
}

// NewMockRequestConsumer creates a new mock instance
func NewMockRequestConsumer(ctrl *gomock.Controller) *MockRequestConsumer {
	mock := &MockRequestConsumer{ctrl: ctrl}
	mock.recorder = &MockRequestConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequestConsumer) EXPECT() *MockRequestConsumerMockRecorder {
	return m.recorder
}

// Deliver mocks base method
func (m *MockRequestConsumer) Deliver(arg0 []byte) []byte {
	ret := m.ctrl.Call(m, "Deliver", arg0)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Deliver indicates an expected call of Deliver
func (mr *MockRequestConsumerMockRecorder) Deliver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deliver", reflect.TypeOf((*MockRequestConsumer)(nil).Deliver), arg0)
}

// StateDigest mocks base method
func (m *MockRequestConsumer) StateDigest() []byte {
	ret := m.ctrl.Call(m, "StateDigest")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// StateDigest indicates an expected call of StateDigest
func (mr *MockRequestConsumerMockRecorder) StateDigest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateDigest", reflect.TypeOf((*MockRequestConsumer)(nil).StateDigest))
}
