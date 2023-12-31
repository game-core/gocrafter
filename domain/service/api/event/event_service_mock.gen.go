// Code generated by MockGen. DO NOT EDIT.
// Source: ./event_service.go

// Package event is a generated GoMock package.
package event

import (
	reflect "reflect"

	event "github.com/game-core/gocrafter/domain/entity/master/event"
	gomock "github.com/golang/mock/gomock"
)

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// GetEventToEntity mocks base method.
func (m *MockEventService) GetEventToEntity(name string) (*event.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventToEntity", name)
	ret0, _ := ret[0].(*event.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventToEntity indicates an expected call of GetEventToEntity.
func (mr *MockEventServiceMockRecorder) GetEventToEntity(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventToEntity", reflect.TypeOf((*MockEventService)(nil).GetEventToEntity), name)
}
