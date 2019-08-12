// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lucas-clemente/quic-go (interfaces: FrameSource)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/lucas-clemente/quic-go/internal/protocol"
	wire "github.com/lucas-clemente/quic-go/internal/wire"
)

// MockFrameSource is a mock of FrameSource interface
type MockFrameSource struct {
	ctrl     *gomock.Controller
	recorder *MockFrameSourceMockRecorder
}

// MockFrameSourceMockRecorder is the mock recorder for MockFrameSource
type MockFrameSourceMockRecorder struct {
	mock *MockFrameSource
}

// NewMockFrameSource creates a new mock instance
func NewMockFrameSource(ctrl *gomock.Controller) *MockFrameSource {
	mock := &MockFrameSource{ctrl: ctrl}
	mock.recorder = &MockFrameSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFrameSource) EXPECT() *MockFrameSourceMockRecorder {
	return m.recorder
}

// AppendControlFrames mocks base method
func (m *MockFrameSource) AppendControlFrames(arg0 []wire.Frame, arg1 protocol.ByteCount) ([]wire.Frame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendControlFrames", arg0, arg1)
	ret0, _ := ret[0].([]wire.Frame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// AppendControlFrames indicates an expected call of AppendControlFrames
func (mr *MockFrameSourceMockRecorder) AppendControlFrames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendControlFrames", reflect.TypeOf((*MockFrameSource)(nil).AppendControlFrames), arg0, arg1)
}

// AppendStreamFrames mocks base method
func (m *MockFrameSource) AppendStreamFrames(arg0 []wire.Frame, arg1 protocol.ByteCount) ([]wire.Frame, protocol.ByteCount) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendStreamFrames", arg0, arg1)
	ret0, _ := ret[0].([]wire.Frame)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// AppendStreamFrames indicates an expected call of AppendStreamFrames
func (mr *MockFrameSourceMockRecorder) AppendStreamFrames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendStreamFrames", reflect.TypeOf((*MockFrameSource)(nil).AppendStreamFrames), arg0, arg1)
}
