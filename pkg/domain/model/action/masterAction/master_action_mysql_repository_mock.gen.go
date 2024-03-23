// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_action_mysql_repository.gen.go

// Package masterAction is a generated GoMock package.
package masterAction

import (
	context "context"
	reflect "reflect"

	enum "github.com/game-core/gocrafter/pkg/domain/enum"
	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterActionMysqlRepository is a mock of MasterActionMysqlRepository interface.
type MockMasterActionMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterActionMysqlRepositoryMockRecorder
}

// MockMasterActionMysqlRepositoryMockRecorder is the mock recorder for MockMasterActionMysqlRepository.
type MockMasterActionMysqlRepositoryMockRecorder struct {
	mock *MockMasterActionMysqlRepository
}

// NewMockMasterActionMysqlRepository creates a new mock instance.
func NewMockMasterActionMysqlRepository(ctrl *gomock.Controller) *MockMasterActionMysqlRepository {
	mock := &MockMasterActionMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterActionMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterActionMysqlRepository) EXPECT() *MockMasterActionMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterActionMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterActionMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterActions) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterActionMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterAction) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockMasterActionMysqlRepository) Find(ctx context.Context, id int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Find), ctx, id)
}

// FindByActionStepType mocks base method.
func (m *MockMasterActionMysqlRepository) FindByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActionStepType", ctx, actionStepType)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActionStepType indicates an expected call of FindByActionStepType.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByActionStepType(ctx, actionStepType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActionStepType", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByActionStepType), ctx, actionStepType)
}

// FindByActionStepTypeAndAnyId mocks base method.
func (m *MockMasterActionMysqlRepository) FindByActionStepTypeAndAnyId(ctx context.Context, actionStepType enum.ActionStepType, anyId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActionStepTypeAndAnyId", ctx, actionStepType, anyId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActionStepTypeAndAnyId indicates an expected call of FindByActionStepTypeAndAnyId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByActionStepTypeAndAnyId(ctx, actionStepType, anyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActionStepTypeAndAnyId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByActionStepTypeAndAnyId), ctx, actionStepType, anyId)
}

// FindByAnyId mocks base method.
func (m *MockMasterActionMysqlRepository) FindByAnyId(ctx context.Context, anyId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAnyId", ctx, anyId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByAnyId indicates an expected call of FindByAnyId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByAnyId(ctx, anyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAnyId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByAnyId), ctx, anyId)
}

// FindByName mocks base method.
func (m *MockMasterActionMysqlRepository) FindByName(ctx context.Context, name string) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindByName), ctx, name)
}

// FindList mocks base method.
func (m *MockMasterActionMysqlRepository) FindList(ctx context.Context) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindList), ctx)
}

// FindListByActionStepType mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByActionStepType", ctx, actionStepType)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByActionStepType indicates an expected call of FindListByActionStepType.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByActionStepType(ctx, actionStepType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByActionStepType", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByActionStepType), ctx, actionStepType)
}

// FindListByActionStepTypeAndAnyId mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByActionStepTypeAndAnyId(ctx context.Context, actionStepType enum.ActionStepType, anyId *int64) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByActionStepTypeAndAnyId", ctx, actionStepType, anyId)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByActionStepTypeAndAnyId indicates an expected call of FindListByActionStepTypeAndAnyId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByActionStepTypeAndAnyId(ctx, actionStepType, anyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByActionStepTypeAndAnyId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByActionStepTypeAndAnyId), ctx, actionStepType, anyId)
}

// FindListByAnyId mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByAnyId(ctx context.Context, anyId *int64) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByAnyId", ctx, anyId)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByAnyId indicates an expected call of FindListByAnyId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByAnyId(ctx, anyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByAnyId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByAnyId), ctx, anyId)
}

// FindListByName mocks base method.
func (m *MockMasterActionMysqlRepository) FindListByName(ctx context.Context, name string) (MasterActions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByName", ctx, name)
	ret0, _ := ret[0].(MasterActions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByName indicates an expected call of FindListByName.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindListByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByName", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindListByName), ctx, name)
}

// FindOrNil mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNil(ctx context.Context, id int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNil), ctx, id)
}

// FindOrNilByActionStepType mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByActionStepType", ctx, actionStepType)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByActionStepType indicates an expected call of FindOrNilByActionStepType.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByActionStepType(ctx, actionStepType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByActionStepType", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByActionStepType), ctx, actionStepType)
}

// FindOrNilByActionStepTypeAndAnyId mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByActionStepTypeAndAnyId(ctx context.Context, actionStepType enum.ActionStepType, anyId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByActionStepTypeAndAnyId", ctx, actionStepType, anyId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByActionStepTypeAndAnyId indicates an expected call of FindOrNilByActionStepTypeAndAnyId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByActionStepTypeAndAnyId(ctx, actionStepType, anyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByActionStepTypeAndAnyId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByActionStepTypeAndAnyId), ctx, actionStepType, anyId)
}

// FindOrNilByAnyId mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByAnyId(ctx context.Context, anyId *int64) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByAnyId", ctx, anyId)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByAnyId indicates an expected call of FindOrNilByAnyId.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByAnyId(ctx, anyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByAnyId", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByAnyId), ctx, anyId)
}

// FindOrNilByName mocks base method.
func (m *MockMasterActionMysqlRepository) FindOrNilByName(ctx context.Context, name string) (*MasterAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByName", ctx, name)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByName indicates an expected call of FindOrNilByName.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) FindOrNilByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByName", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).FindOrNilByName), ctx, name)
}

// Update mocks base method.
func (m_2 *MockMasterActionMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterActionMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterActionMysqlRepository)(nil).Update), ctx, tx, m)
}