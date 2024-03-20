// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_idle_bonus_item_repository.gen.go

// Package masterIdleBonusItem is a generated GoMock package.
package masterIdleBonusItem

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterIdleBonusItemRepository is a mock of MasterIdleBonusItemRepository interface.
type MockMasterIdleBonusItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterIdleBonusItemRepositoryMockRecorder
}

// MockMasterIdleBonusItemRepositoryMockRecorder is the mock recorder for MockMasterIdleBonusItemRepository.
type MockMasterIdleBonusItemRepositoryMockRecorder struct {
	mock *MockMasterIdleBonusItemRepository
}

// NewMockMasterIdleBonusItemRepository creates a new mock instance.
func NewMockMasterIdleBonusItemRepository(ctrl *gomock.Controller) *MockMasterIdleBonusItemRepository {
	mock := &MockMasterIdleBonusItemRepository{ctrl: ctrl}
	mock.recorder = &MockMasterIdleBonusItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterIdleBonusItemRepository) EXPECT() *MockMasterIdleBonusItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMasterIdleBonusItemRepository) Create(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusItem) (*MasterIdleBonusItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockMasterIdleBonusItemRepository) CreateList(ctx context.Context, tx *gorm.DB, ms MasterIdleBonusItems) (MasterIdleBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(MasterIdleBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockMasterIdleBonusItemRepository) Delete(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusItem) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).Delete), ctx, tx, m)
}

// FindOrNilByMasterIdleBonusScheduleId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindOrNilByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterIdleBonusScheduleId", ctx, masterIdleBonusScheduleId)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterIdleBonusScheduleId indicates an expected call of FindOrNilByMasterIdleBonusScheduleId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindOrNilByMasterIdleBonusScheduleId(ctx, masterIdleBonusScheduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterIdleBonusScheduleId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindOrNilByMasterIdleBonusScheduleId), ctx, masterIdleBonusScheduleId)
}

// FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId, masterItemId int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId", ctx, masterIdleBonusScheduleId, masterItemId)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId indicates an expected call of FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId(ctx, masterIdleBonusScheduleId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindOrNilByMasterIdleBonusScheduleIdAndMasterItemId), ctx, masterIdleBonusScheduleId, masterItemId)
}

// FindOrNilByMasterItemId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindOrNilByMasterItemId(ctx context.Context, masterItemId int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterItemId", ctx, masterItemId)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterItemId indicates an expected call of FindOrNilByMasterItemId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindOrNilByMasterItemId(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterItemId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindOrNilByMasterItemId), ctx, masterItemId)
}

// Find mocks base method.
func (m *MockMasterIdleBonusItemRepository) Find(ctx context.Context, id int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, id)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) Find(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).Find), ctx, id)
}

// FindByMasterIdleBonusScheduleId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterIdleBonusScheduleId", ctx, masterIdleBonusScheduleId)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterIdleBonusScheduleId indicates an expected call of FindByMasterIdleBonusScheduleId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindByMasterIdleBonusScheduleId(ctx, masterIdleBonusScheduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterIdleBonusScheduleId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindByMasterIdleBonusScheduleId), ctx, masterIdleBonusScheduleId)
}

// FindByMasterIdleBonusScheduleIdAndMasterItemId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId, masterItemId int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterIdleBonusScheduleIdAndMasterItemId", ctx, masterIdleBonusScheduleId, masterItemId)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterIdleBonusScheduleIdAndMasterItemId indicates an expected call of FindByMasterIdleBonusScheduleIdAndMasterItemId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindByMasterIdleBonusScheduleIdAndMasterItemId(ctx, masterIdleBonusScheduleId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterIdleBonusScheduleIdAndMasterItemId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindByMasterIdleBonusScheduleIdAndMasterItemId), ctx, masterIdleBonusScheduleId, masterItemId)
}

// FindByMasterItemId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindByMasterItemId(ctx context.Context, masterItemId int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterItemId", ctx, masterItemId)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterItemId indicates an expected call of FindByMasterItemId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindByMasterItemId(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterItemId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindByMasterItemId), ctx, masterItemId)
}

// FindList mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindList(ctx context.Context) (MasterIdleBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(MasterIdleBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindList), ctx)
}

// FindListByMasterIdleBonusScheduleId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindListByMasterIdleBonusScheduleId(ctx context.Context, masterIdleBonusScheduleId int64) (MasterIdleBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterIdleBonusScheduleId", ctx, masterIdleBonusScheduleId)
	ret0, _ := ret[0].(MasterIdleBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterIdleBonusScheduleId indicates an expected call of FindListByMasterIdleBonusScheduleId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindListByMasterIdleBonusScheduleId(ctx, masterIdleBonusScheduleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterIdleBonusScheduleId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindListByMasterIdleBonusScheduleId), ctx, masterIdleBonusScheduleId)
}

// FindListByMasterIdleBonusScheduleIdAndMasterItemId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindListByMasterIdleBonusScheduleIdAndMasterItemId(ctx context.Context, masterIdleBonusScheduleId, masterItemId int64) (MasterIdleBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterIdleBonusScheduleIdAndMasterItemId", ctx, masterIdleBonusScheduleId, masterItemId)
	ret0, _ := ret[0].(MasterIdleBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterIdleBonusScheduleIdAndMasterItemId indicates an expected call of FindListByMasterIdleBonusScheduleIdAndMasterItemId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindListByMasterIdleBonusScheduleIdAndMasterItemId(ctx, masterIdleBonusScheduleId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterIdleBonusScheduleIdAndMasterItemId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindListByMasterIdleBonusScheduleIdAndMasterItemId), ctx, masterIdleBonusScheduleId, masterItemId)
}

// FindListByMasterItemId mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindListByMasterItemId(ctx context.Context, masterItemId int64) (MasterIdleBonusItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterItemId", ctx, masterItemId)
	ret0, _ := ret[0].(MasterIdleBonusItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterItemId indicates an expected call of FindListByMasterItemId.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindListByMasterItemId(ctx, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterItemId", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindListByMasterItemId), ctx, masterItemId)
}

// FindOrNil mocks base method.
func (m *MockMasterIdleBonusItemRepository) FindOrNil(ctx context.Context, id int64) (*MasterIdleBonusItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, id)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) FindOrNil(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).FindOrNil), ctx, id)
}

// Update mocks base method.
func (m_2 *MockMasterIdleBonusItemRepository) Update(ctx context.Context, tx *gorm.DB, m *MasterIdleBonusItem) (*MasterIdleBonusItem, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*MasterIdleBonusItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMasterIdleBonusItemRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMasterIdleBonusItemRepository)(nil).Update), ctx, tx, m)
}
