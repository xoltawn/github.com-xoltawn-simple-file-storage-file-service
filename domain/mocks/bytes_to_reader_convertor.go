// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xoltawn/simple-file-storage-file-service/domain (interfaces: BytesToReaderConvertor)

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBytesToReaderConvertor is a mock of BytesToReaderConvertor interface.
type MockBytesToReaderConvertor struct {
	ctrl     *gomock.Controller
	recorder *MockBytesToReaderConvertorMockRecorder
}

// MockBytesToReaderConvertorMockRecorder is the mock recorder for MockBytesToReaderConvertor.
type MockBytesToReaderConvertorMockRecorder struct {
	mock *MockBytesToReaderConvertor
}

// NewMockBytesToReaderConvertor creates a new mock instance.
func NewMockBytesToReaderConvertor(ctrl *gomock.Controller) *MockBytesToReaderConvertor {
	mock := &MockBytesToReaderConvertor{ctrl: ctrl}
	mock.recorder = &MockBytesToReaderConvertorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBytesToReaderConvertor) EXPECT() *MockBytesToReaderConvertorMockRecorder {
	return m.recorder
}

// Convert mocks base method.
func (m *MockBytesToReaderConvertor) Convert(arg0 []byte) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Convert", arg0)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Convert indicates an expected call of Convert.
func (mr *MockBytesToReaderConvertorMockRecorder) Convert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Convert", reflect.TypeOf((*MockBytesToReaderConvertor)(nil).Convert), arg0)
}
