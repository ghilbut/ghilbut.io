// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../../common/proxy.go

// package mock_common
package mock

import (
	. "github.com/ghilbut/ygg.go/common"
	gomock "github.com/golang/mock/gomock"
)

// Mock of CtrlProxyDelegate interface
type MockCtrlProxyDelegate struct {
	ctrl     *gomock.Controller
	recorder *_MockCtrlProxyDelegateRecorder
}

// Recorder for MockCtrlProxyDelegate (not exported)
type _MockCtrlProxyDelegateRecorder struct {
	mock *MockCtrlProxyDelegate
}

func NewMockCtrlProxyDelegate(ctrl *gomock.Controller) *MockCtrlProxyDelegate {
	mock := &MockCtrlProxyDelegate{ctrl: ctrl}
	mock.recorder = &_MockCtrlProxyDelegateRecorder{mock}
	return mock
}

func (_m *MockCtrlProxyDelegate) EXPECT() *_MockCtrlProxyDelegateRecorder {
	return _m.recorder
}

func (_m *MockCtrlProxyDelegate) OnCtrlText(proxy *CtrlProxy, text string) {
	_m.ctrl.Call(_m, "OnCtrlText", proxy, text)
}

func (_mr *_MockCtrlProxyDelegateRecorder) OnCtrlText(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnCtrlText", arg0, arg1)
}

func (_m *MockCtrlProxyDelegate) OnCtrlBinary(proxy *CtrlProxy, bytes []byte) {
	_m.ctrl.Call(_m, "OnCtrlBinary", proxy, bytes)
}

func (_mr *_MockCtrlProxyDelegateRecorder) OnCtrlBinary(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnCtrlBinary", arg0, arg1)
}

func (_m *MockCtrlProxyDelegate) OnCtrlClosed(proxy *CtrlProxy) {
	_m.ctrl.Call(_m, "OnCtrlClosed", proxy)
}

func (_mr *_MockCtrlProxyDelegateRecorder) OnCtrlClosed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnCtrlClosed", arg0)
}

// Mock of TargetProxyDelegate interface
type MockTargetProxyDelegate struct {
	ctrl     *gomock.Controller
	recorder *_MockTargetProxyDelegateRecorder
}

// Recorder for MockTargetProxyDelegate (not exported)
type _MockTargetProxyDelegateRecorder struct {
	mock *MockTargetProxyDelegate
}

func NewMockTargetProxyDelegate(ctrl *gomock.Controller) *MockTargetProxyDelegate {
	mock := &MockTargetProxyDelegate{ctrl: ctrl}
	mock.recorder = &_MockTargetProxyDelegateRecorder{mock}
	return mock
}

func (_m *MockTargetProxyDelegate) EXPECT() *_MockTargetProxyDelegateRecorder {
	return _m.recorder
}

func (_m *MockTargetProxyDelegate) OnTargetText(proxy *TargetProxy, text string) {
	_m.ctrl.Call(_m, "OnTargetText", proxy, text)
}

func (_mr *_MockTargetProxyDelegateRecorder) OnTargetText(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnTargetText", arg0, arg1)
}

func (_m *MockTargetProxyDelegate) OnTargetBinary(proxy *TargetProxy, bytes []byte) {
	_m.ctrl.Call(_m, "OnTargetBinary", proxy, bytes)
}

func (_mr *_MockTargetProxyDelegateRecorder) OnTargetBinary(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnTargetBinary", arg0, arg1)
}

func (_m *MockTargetProxyDelegate) OnTargetClosed(proxy *TargetProxy) {
	_m.ctrl.Call(_m, "OnTargetClosed", proxy)
}

func (_mr *_MockTargetProxyDelegateRecorder) OnTargetClosed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OnTargetClosed", arg0)
}