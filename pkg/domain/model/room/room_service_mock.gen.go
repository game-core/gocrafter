// Code generated by MockGen. DO NOT EDIT.
// Source: ./room_service.go

// Package room is a generated GoMock package.
package room

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockRoomService is a mock of RoomService interface.
type MockRoomService struct {
	ctrl     *gomock.Controller
	recorder *MockRoomServiceMockRecorder
}

// MockRoomServiceMockRecorder is the mock recorder for MockRoomService.
type MockRoomServiceMockRecorder struct {
	mock *MockRoomService
}

// NewMockRoomService creates a new mock instance.
func NewMockRoomService(ctrl *gomock.Controller) *MockRoomService {
	mock := &MockRoomService{ctrl: ctrl}
	mock.recorder = &MockRoomServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomService) EXPECT() *MockRoomServiceMockRecorder {
	return m.recorder
}

// CheckIn mocks base method.
func (m *MockRoomService) CheckIn(ctx context.Context, tx *gorm.DB, req *RoomCheckInRequest) (*RoomCheckInResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIn", ctx, tx, req)
	ret0, _ := ret[0].(*RoomCheckInResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIn indicates an expected call of CheckIn.
func (mr *MockRoomServiceMockRecorder) CheckIn(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIn", reflect.TypeOf((*MockRoomService)(nil).CheckIn), ctx, tx, req)
}

// CheckOut mocks base method.
func (m *MockRoomService) CheckOut(ctx context.Context, tx *gorm.DB, req *RoomCheckOutRequest) (*RoomCheckOutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOut", ctx, tx, req)
	ret0, _ := ret[0].(*RoomCheckOutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOut indicates an expected call of CheckOut.
func (mr *MockRoomServiceMockRecorder) CheckOut(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOut", reflect.TypeOf((*MockRoomService)(nil).CheckOut), ctx, tx, req)
}

// Create mocks base method.
func (m *MockRoomService) Create(ctx context.Context, tx *gorm.DB, req *RoomCreateRequest) (*RoomCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tx, req)
	ret0, _ := ret[0].(*RoomCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoomServiceMockRecorder) Create(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoomService)(nil).Create), ctx, tx, req)
}

// Delete mocks base method.
func (m *MockRoomService) Delete(ctx context.Context, tx *gorm.DB, req *RoomDeleteRequest) (*RoomDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, tx, req)
	ret0, _ := ret[0].(*RoomDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRoomServiceMockRecorder) Delete(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoomService)(nil).Delete), ctx, tx, req)
}

// Search mocks base method.
func (m *MockRoomService) Search(ctx context.Context, req *RoomSearchRequest) (*RoomSearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, req)
	ret0, _ := ret[0].(*RoomSearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRoomServiceMockRecorder) Search(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRoomService)(nil).Search), ctx, req)
}
