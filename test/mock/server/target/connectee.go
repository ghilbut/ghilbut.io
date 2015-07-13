// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../../server/target/connectee.go

package mock_target

import (
	. "github.com/ghilbut/ygg.go/common"
	. "github.com/ghilbut/ygg.go/server/target"
	gomock "github.com/golang/mock/gomock"
)

// Mock of Connectee interface
type MockConnectee struct {
	ctrl     *gomock.Controller
	recorder *_MockConnecteeRecorder
}

// Recorder for MockConnectee (not exported)
type _MockConnecteeRecorder struct {
	mock *MockConnectee
}

func NewMockConnectee(ctrl *gomock.Controller) *MockConnectee {
	mock := &MockConnectee{ctrl: ctrl}
	mock.recorder = &_MockConnecteeRecorder{mock}
	return mock
}

func (_m *MockConnectee) EXPECT() *_MockConnecteeRecorder {
	return _m.recorder
}

func (_m *MockConnectee) Start(delegate ConnecteeDelegate) {
	_m.ctrl.Call(_m, "Start", delegate)
}

func (_mr *_MockConnecteeRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockConnectee) Stop() {
	_m.ctrl.Call(_m, "Stop")
}

func (_mr *_MockConnecteeRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

func (_m *MockConnectee) Register(endpoint string) bool {
	ret := _m.ctrl.Call(_m, "Register", endpoint)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockConnecteeRecorder) Register(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Register", arg0)
}

func (_m *MockConnectee) Unregister(endpoint string) {
	_m.ctrl.Call(_m, "Unregister", endpoint)
}

func (_mr *_MockConnecteeRecorder) Unregister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unregister", arg0)
}

func (_m *MockConnectee) HasEndpoint(endpoint string) bool {
	ret := _m.ctrl.Call(_m, "HasEndpoint", endpoint)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockConnecteeRecorder) HasEndpoint(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasEndpoint", arg0)
}

// Mock of ConnecteeDelegate interface
type MockConnecteeDelegate struct {
	ctrl     *gomock.Controller
	recorder *_MockConnecteeDelegateRecorder
}

// Recorder for MockConnecteeDelegate (not exported)
type _MockConnecteeDelegateRecorder struct {
	mock *MockConnecteeDelegate
}

func NewMockConnecteeDelegate(ctrl *gomock.Controller) *MockConnecteeDelegate {
	mock := &MockConnecteeDelegate{ctrl: ctrl}
	mock.recorder = &_MockConnecteeDelegateRecorder{mock}
	return mock
}

func (_m *MockConnecteeDelegate) EXPECT() *_MockConnecteeDelegateRecorder {
	return _m.recorder
}

func (_m *MockConnecteeDelegate) OnCtrlConnected(conn Connection) {
	_m.ctrl.Call(_m, "OnCtrlConnected", conn)
}

func (_mr *_MockConnecteeDelegateRecorder) OnCtrlConnected(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnCtrlConnected", arg0)
}

func (_m *MockConnecteeDelegate) OnTargetConnected(conn Connection) {
	_m.ctrl.Call(_m, "OnTargetConnected", conn)
}

func (_mr *_MockConnecteeDelegateRecorder) OnTargetConnected(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnTargetConnected", arg0)
}
