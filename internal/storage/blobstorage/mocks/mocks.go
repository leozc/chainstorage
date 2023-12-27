// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coinbase/chainstorage/internal/storage/blobstorage (interfaces: BlobStorage)
//
// Generated by this command:
//
//	mockgen -destination internal/storage/blobstorage/mocks/mocks.go -package blobstoragemocks github.com/coinbase/chainstorage/internal/storage/blobstorage BlobStorage
//

// Package blobstoragemocks is a generated GoMock package.
package blobstoragemocks

import (
	context "context"
	reflect "reflect"

	chainstorage "github.com/coinbase/chainstorage/protos/coinbase/chainstorage"
	gomock "go.uber.org/mock/gomock"
)

// MockBlobStorage is a mock of BlobStorage interface.
type MockBlobStorage struct {
	ctrl     *gomock.Controller
	recorder *MockBlobStorageMockRecorder
}

// MockBlobStorageMockRecorder is the mock recorder for MockBlobStorage.
type MockBlobStorageMockRecorder struct {
	mock *MockBlobStorage
}

// NewMockBlobStorage creates a new mock instance.
func NewMockBlobStorage(ctrl *gomock.Controller) *MockBlobStorage {
	mock := &MockBlobStorage{ctrl: ctrl}
	mock.recorder = &MockBlobStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlobStorage) EXPECT() *MockBlobStorageMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockBlobStorage) Download(arg0 context.Context, arg1 *chainstorage.BlockMetadata) (*chainstorage.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(*chainstorage.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockBlobStorageMockRecorder) Download(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockBlobStorage)(nil).Download), arg0, arg1)
}

// Upload mocks base method.
func (m *MockBlobStorage) Upload(arg0 context.Context, arg1 *chainstorage.Block, arg2 chainstorage.Compression) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockBlobStorageMockRecorder) Upload(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockBlobStorage)(nil).Upload), arg0, arg1, arg2)
}
