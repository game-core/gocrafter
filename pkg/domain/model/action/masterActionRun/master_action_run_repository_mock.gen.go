// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_action_run_repository.gen.go

// Package masterActionRun is a generated GoMock package.
package masterActionRun

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterActionRunMysqlRepository is a mock of MasterActionRunMysqlRepository interface.
type MockMasterActionRunMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterActionRunMysqlRepositoryMockRecorder
}

// MockMasterActionRunMysqlRepositoryMockRecorder is the mock recorder for MockMasterActionRunMysqlRepository.
type MockMasterActionRunMysqlRepositoryMockRecorder struct {
	mock *MockMasterActionRunMysqlRepository
}

// NewMockMasterActionRunMysqlRepository creates a new mock instance.
func NewMockMasterActionRunMysqlRepository(ctrl *gomock.Controller) *MockMasterActionRunMysqlRepository {
	mock := &MockMasterActionRunMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterActionRunMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterActionRunMysqlRepository) EXPECT() *MockMasterActionRunMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterActionRunMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterActionRun) (*MasterActionRun, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterActionRunMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionRuns) (MasterActionRuns, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterActionRuns)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterActionRunMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterActionRun) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterActionRunMysqlRepository) Find(ctx context.Context, id int64) (*MasterActionRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).Find), ctx, id)
}

// FindByActionId mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindByActionId(ctx context.Context, actionId int64) (*MasterActionRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActionId", ctx, actionId)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActionId indicates an expected call of FindByActionId.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindByActionId(ctx, actionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActionId", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindByActionId), ctx, actionId)
}

// FindByName mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindByName(ctx context.Context, name string) (*MasterActionRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindByName), ctx, name)
}

// FindList mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindList(ctx context.Context) (MasterActionRuns, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterActionRuns)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindList), ctx)
}

// FindListByActionId mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindListByActionId(ctx context.Context, actionId int64) (MasterActionRuns, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByActionId", ctx, actionId)
	ret0, _ := ret[0].(MasterActionRuns)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByActionId indicates an expected call of FindListByActionId.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindListByActionId(ctx, actionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByActionId", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindListByActionId), ctx, actionId)
}

// FindListByName mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindListByName(ctx context.Context, name string) (MasterActionRuns, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByName", ctx, name)
	ret0, _ := ret[0].(MasterActionRuns)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByName indicates an expected call of FindListByName.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindListByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByName", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindListByName), ctx, name)
}

// FindOrNil mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindOrNil(ctx context.Context, id int64) (*MasterActionRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindOrNil), ctx, id)
}

// FindOrNilByActionId mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindOrNilByActionId(ctx context.Context, actionId int64) (*MasterActionRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByActionId", ctx, actionId)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByActionId indicates an expected call of FindOrNilByActionId.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindOrNilByActionId(ctx, actionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByActionId", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindOrNilByActionId), ctx, actionId)
}

// FindOrNilByName mocks base method.
func (m *MockMasterActionRunMysqlRepository) FindOrNilByName(ctx context.Context, name string) (*MasterActionRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByName", ctx, name)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByName indicates an expected call of FindOrNilByName.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) FindOrNilByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByName", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).FindOrNilByName), ctx, name)
}

// Update mocks base method.
func (m_2 *MockMasterActionRunMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterActionRun) (*MasterActionRun, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterActionRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterActionRunMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterActionRunMysqlRepository)(nil).Update), ctx, tx, m)
}
