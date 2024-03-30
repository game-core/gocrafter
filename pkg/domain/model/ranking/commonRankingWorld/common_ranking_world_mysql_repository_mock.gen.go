// Code generated by MockGen. DO NOT EDIT.
// Source: ./common_ranking_world_mysql_repository.gen.go

// Package commonRankingWorld is a generated GoMock package.
package commonRankingWorld

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockCommonRankingWorldMysqlRepository is a mock of CommonRankingWorldMysqlRepository interface.
type MockCommonRankingWorldMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommonRankingWorldMysqlRepositoryMockRecorder
}

// MockCommonRankingWorldMysqlRepositoryMockRecorder is the mock recorder for MockCommonRankingWorldMysqlRepository.
type MockCommonRankingWorldMysqlRepositoryMockRecorder struct {
	mock *MockCommonRankingWorldMysqlRepository
}

// NewMockCommonRankingWorldMysqlRepository creates a new mock instance.
func NewMockCommonRankingWorldMysqlRepository(ctrl *gomock.Controller) *MockCommonRankingWorldMysqlRepository {
	mock := &MockCommonRankingWorldMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockCommonRankingWorldMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonRankingWorldMysqlRepository) EXPECT() *MockCommonRankingWorldMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockCommonRankingWorldMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *CommonRankingWorld) (*CommonRankingWorld, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*CommonRankingWorld)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms CommonRankingWorlds) (CommonRankingWorlds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(CommonRankingWorlds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockCommonRankingWorldMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *CommonRankingWorld) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) Find(ctx context.Context, masterRankingId int64, userId string) (*CommonRankingWorld, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterRankingId, userId)
	ret0, _ := ret[0].(*CommonRankingWorld)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) Find(ctx, masterRankingId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).Find), ctx, masterRankingId, userId)
}

// FindByMasterRankingId mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) FindByMasterRankingId(ctx context.Context, masterRankingId int64) (*CommonRankingWorld, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterRankingId", ctx, masterRankingId)
	ret0, _ := ret[0].(*CommonRankingWorld)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterRankingId indicates an expected call of FindByMasterRankingId.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) FindByMasterRankingId(ctx, masterRankingId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterRankingId", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).FindByMasterRankingId), ctx, masterRankingId)
}

// FindList mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) FindList(ctx context.Context) (CommonRankingWorlds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(CommonRankingWorlds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).FindList), ctx)
}

// FindListByMasterRankingId mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) FindListByMasterRankingId(ctx context.Context, masterRankingId int64) (CommonRankingWorlds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterRankingId", ctx, masterRankingId)
	ret0, _ := ret[0].(CommonRankingWorlds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterRankingId indicates an expected call of FindListByMasterRankingId.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) FindListByMasterRankingId(ctx, masterRankingId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterRankingId", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).FindListByMasterRankingId), ctx, masterRankingId)
}

// FindOrNil mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) FindOrNil(ctx context.Context, masterRankingId int64, userId string) (*CommonRankingWorld, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterRankingId, userId)
	ret0, _ := ret[0].(*CommonRankingWorld)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) FindOrNil(ctx, masterRankingId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).FindOrNil), ctx, masterRankingId, userId)
}

// FindOrNilByMasterRankingId mocks base method.
func (m *MockCommonRankingWorldMysqlRepository) FindOrNilByMasterRankingId(ctx context.Context, masterRankingId int64) (*CommonRankingWorld, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterRankingId", ctx, masterRankingId)
	ret0, _ := ret[0].(*CommonRankingWorld)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterRankingId indicates an expected call of FindOrNilByMasterRankingId.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) FindOrNilByMasterRankingId(ctx, masterRankingId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterRankingId", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).FindOrNilByMasterRankingId), ctx, masterRankingId)
}

// Update mocks base method.
func (m_2 *MockCommonRankingWorldMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *CommonRankingWorld) (*CommonRankingWorld, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*CommonRankingWorld)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommonRankingWorldMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommonRankingWorldMysqlRepository)(nil).Update), ctx, tx, m)
}