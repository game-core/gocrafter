// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_login_bonus_event_mysql_repository.gen.go

// Package masterLoginBonusEvent is a generated GoMock package.
package masterLoginBonusEvent

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterLoginBonusEventMysqlRepository is a mock of MasterLoginBonusEventMysqlRepository interface.
type MockMasterLoginBonusEventMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterLoginBonusEventMysqlRepositoryMockRecorder
}

// MockMasterLoginBonusEventMysqlRepositoryMockRecorder is the mock recorder for MockMasterLoginBonusEventMysqlRepository.
type MockMasterLoginBonusEventMysqlRepositoryMockRecorder struct {
	mock *MockMasterLoginBonusEventMysqlRepository
}

// NewMockMasterLoginBonusEventMysqlRepository creates a new mock instance.
func NewMockMasterLoginBonusEventMysqlRepository(ctrl *gomock.Controller) *MockMasterLoginBonusEventMysqlRepository {
	mock := &MockMasterLoginBonusEventMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterLoginBonusEventMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterLoginBonusEventMysqlRepository) EXPECT() *MockMasterLoginBonusEventMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterLoginBonusEventMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusEvent) (*MasterLoginBonusEvent, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonusEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterLoginBonusEventMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonusEvents) (MasterLoginBonusEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterLoginBonusEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterLoginBonusEventMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusEvent) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterLoginBonusEventMysqlRepository) Find(ctx context.Context, id int64) (*MasterLoginBonusEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonusEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).Find), ctx, id)
}

// FindList mocks base method.
func (m *MockMasterLoginBonusEventMysqlRepository) FindList(ctx context.Context) (MasterLoginBonusEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterLoginBonusEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).FindList), ctx)
}

// FindOrNil mocks base method.
func (m *MockMasterLoginBonusEventMysqlRepository) FindOrNil(ctx context.Context, id int64) (*MasterLoginBonusEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterLoginBonusEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).FindOrNil), ctx, id)
}

// Update mocks base method.
func (m_2 *MockMasterLoginBonusEventMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonusEvent) (*MasterLoginBonusEvent, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterLoginBonusEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterLoginBonusEventMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterLoginBonusEventMysqlRepository)(nil).Update), ctx, tx, m)
}