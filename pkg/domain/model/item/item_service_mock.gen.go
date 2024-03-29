// Code generated by MockGen. DO NOT EDIT.
// Source: ./item_service.go

// Package item is a generated GoMock package.
package item

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockItemService is a mock of ItemService interface.
type MockItemService struct {
	ctrl     *gomock.Controller
	recorder *MockItemServiceMockRecorder
}

// MockItemServiceMockRecorder is the mock recorder for MockItemService.
type MockItemServiceMockRecorder struct {
	mock *MockItemService
}

// NewMockItemService creates a new mock instance.
func NewMockItemService(ctrl *gomock.Controller) *MockItemService {
	mock := &MockItemService{ctrl: ctrl}
	mock.recorder = &MockItemServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemService) EXPECT() *MockItemServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockItemService) Create(ctx context.Context, tx *gorm.DB, req *ItemCreateRequest) (*ItemCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tx, req)
	ret0, _ := ret[0].(*ItemCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockItemServiceMockRecorder) Create(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemService)(nil).Create), ctx, tx, req)
}

// Receive mocks base method.
func (m *MockItemService) Receive(ctx context.Context, tx *gorm.DB, req *ItemReceiveRequest) (*ItemReceiveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", ctx, tx, req)
	ret0, _ := ret[0].(*ItemReceiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockItemServiceMockRecorder) Receive(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockItemService)(nil).Receive), ctx, tx, req)
}
