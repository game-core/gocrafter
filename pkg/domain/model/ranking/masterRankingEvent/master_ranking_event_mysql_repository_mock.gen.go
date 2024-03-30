// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_ranking_event_mysql_repository.gen.go

// Package masterRankingEvent is a generated GoMock package.
package masterRankingEvent

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterRankingEventMysqlRepository is a mock of MasterRankingEventMysqlRepository interface.
type MockMasterRankingEventMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterRankingEventMysqlRepositoryMockRecorder
}

// MockMasterRankingEventMysqlRepositoryMockRecorder is the mock recorder for MockMasterRankingEventMysqlRepository.
type MockMasterRankingEventMysqlRepositoryMockRecorder struct {
	mock *MockMasterRankingEventMysqlRepository
}

// NewMockMasterRankingEventMysqlRepository creates a new mock instance.
func NewMockMasterRankingEventMysqlRepository(ctrl *gomock.Controller) *MockMasterRankingEventMysqlRepository {
	mock := &MockMasterRankingEventMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterRankingEventMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterRankingEventMysqlRepository) EXPECT() *MockMasterRankingEventMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterRankingEventMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterRankingEvent) (*MasterRankingEvent, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterRankingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterRankingEventMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterRankingEvents) (MasterRankingEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterRankingEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterRankingEventMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterRankingEvent) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterRankingEventMysqlRepository) Find(ctx context.Context, id int64) (*MasterRankingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterRankingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).Find), ctx, id)
}

// FindList mocks base method.
func (m *MockMasterRankingEventMysqlRepository) FindList(ctx context.Context) (MasterRankingEvents, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterRankingEvents)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).FindList), ctx)
}

// FindOrNil mocks base method.
func (m *MockMasterRankingEventMysqlRepository) FindOrNil(ctx context.Context, id int64) (*MasterRankingEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterRankingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).FindOrNil), ctx, id)
}

// Update mocks base method.
func (m_2 *MockMasterRankingEventMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterRankingEvent) (*MasterRankingEvent, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterRankingEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterRankingEventMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterRankingEventMysqlRepository)(nil).Update), ctx, tx, m)
}