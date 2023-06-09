// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xoltawn/simple-file-storage-file-service/domain (interfaces: FileDownloader)

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/xoltawn/simple-file-storage-file-service/domain"
)

// MockFileDownloader is a mock of FileDownloader interface.
type MockFileDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockFileDownloaderMockRecorder
}

// MockFileDownloaderMockRecorder is the mock recorder for MockFileDownloader.
type MockFileDownloaderMockRecorder struct {
	mock *MockFileDownloader
}

// NewMockFileDownloader creates a new mock instance.
func NewMockFileDownloader(ctrl *gomock.Controller) *MockFileDownloader {
	mock := &MockFileDownloader{ctrl: ctrl}
	mock.recorder = &MockFileDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileDownloader) EXPECT() *MockFileDownloaderMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockFileDownloader) Download(arg0 *domain.FileWithBytes) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockFileDownloaderMockRecorder) Download(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockFileDownloader)(nil).Download), arg0)
}
