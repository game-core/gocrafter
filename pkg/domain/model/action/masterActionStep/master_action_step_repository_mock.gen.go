// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_action_step_repository.gen.go

// Package masterActionStep is a generated GoMock package.
package masterActionStep

import (
	context "context"
	reflect "reflect"

	enum "github.com/game-core/gocrafter/pkg/domain/enum"
	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterActionStepRepository is a mock of MasterActionStepRepository interface.
type MockMasterActionStepRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterActionStepRepositoryMockRecorder
}

// MockMasterActionStepRepositoryMockRecorder is the mock recorder for MockMasterActionStepRepository.
type MockMasterActionStepRepositoryMockRecorder struct {
	mock *MockMasterActionStepRepository
}

// NewMockMasterActionStepRepository creates a new mock instance.
func NewMockMasterActionStepRepository(ctrl *gomock.Controller) *MockMasterActionStepRepository {
	mock := &MockMasterActionStepRepository{ctrl: ctrl}
	mock.recorder = &MockMasterActionStepRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterActionStepRepository) EXPECT() *MockMasterActionStepRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterActionStepRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterActionStep) (*MasterActionStep, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterActionStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterActionStepRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterActionStepRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterActionStepRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionSteps) (MasterActionSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterActionSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterActionStepRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterActionStepRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterActionStepRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterActionStep) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterActionStepRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterActionStepRepository)(nil).Delete), ctx, tx, m)
}

// FindOrNilByActionStepType mocks base method.
func (m *MockMasterActionStepRepository) FindOrNilByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterActionStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByActionStepType", ctx, actionStepType)
	ret0, _ := ret[0].(*MasterActionStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByActionStepType indicates an expected call of FindOrNilByActionStepType.
func (mr *MockMasterActionStepRepositoryMockRecorder) FindOrNilByActionStepType(ctx, actionStepType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByActionStepType", reflect.TypeOf((*MockMasterActionStepRepository)(nil).FindOrNilByActionStepType), ctx, actionStepType)
}

// Find mocks base method.
func (m *MockMasterActionStepRepository) Find(ctx context.Context, id int64) (*MasterActionStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterActionStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterActionStepRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterActionStepRepository)(nil).Find), ctx, id)
}

// FindByActionStepType mocks base method.
func (m *MockMasterActionStepRepository) FindByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (*MasterActionStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByActionStepType", ctx, actionStepType)
	ret0, _ := ret[0].(*MasterActionStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByActionStepType indicates an expected call of FindByActionStepType.
func (mr *MockMasterActionStepRepositoryMockRecorder) FindByActionStepType(ctx, actionStepType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByActionStepType", reflect.TypeOf((*MockMasterActionStepRepository)(nil).FindByActionStepType), ctx, actionStepType)
}

// FindList mocks base method.
func (m *MockMasterActionStepRepository) FindList(ctx context.Context) (MasterActionSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterActionSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterActionStepRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterActionStepRepository)(nil).FindList), ctx)
}

// FindListByActionStepType mocks base method.
func (m *MockMasterActionStepRepository) FindListByActionStepType(ctx context.Context, actionStepType enum.ActionStepType) (MasterActionSteps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByActionStepType", ctx, actionStepType)
	ret0, _ := ret[0].(MasterActionSteps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByActionStepType indicates an expected call of FindListByActionStepType.
func (mr *MockMasterActionStepRepositoryMockRecorder) FindListByActionStepType(ctx, actionStepType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByActionStepType", reflect.TypeOf((*MockMasterActionStepRepository)(nil).FindListByActionStepType), ctx, actionStepType)
}

// FindOrNil mocks base method.
func (m *MockMasterActionStepRepository) FindOrNil(ctx context.Context, id int64) (*MasterActionStep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterActionStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterActionStepRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterActionStepRepository)(nil).FindOrNil), ctx, id)
}

// Update mocks base method.
func (m_2 *MockMasterActionStepRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterActionStep) (*MasterActionStep, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterActionStep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterActionStepRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterActionStepRepository)(nil).Update), ctx, tx, m)
}
