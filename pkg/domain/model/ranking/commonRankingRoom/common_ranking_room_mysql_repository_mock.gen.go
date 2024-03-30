// Code generated by MockGen. DO NOT EDIT.
// Source: ./common_ranking_room_mysql_repository.gen.go

// Package commonRankingRoom is a generated GoMock package.
package commonRankingRoom

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockCommonRankingRoomMysqlRepository is a mock of CommonRankingRoomMysqlRepository interface.
type MockCommonRankingRoomMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommonRankingRoomMysqlRepositoryMockRecorder
}

// MockCommonRankingRoomMysqlRepositoryMockRecorder is the mock recorder for MockCommonRankingRoomMysqlRepository.
type MockCommonRankingRoomMysqlRepositoryMockRecorder struct {
	mock *MockCommonRankingRoomMysqlRepository
}

// NewMockCommonRankingRoomMysqlRepository creates a new mock instance.
func NewMockCommonRankingRoomMysqlRepository(ctrl *gomock.Controller) *MockCommonRankingRoomMysqlRepository {
	mock := &MockCommonRankingRoomMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockCommonRankingRoomMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonRankingRoomMysqlRepository) EXPECT() *MockCommonRankingRoomMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockCommonRankingRoomMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *CommonRankingRoom) (*CommonRankingRoom, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*CommonRankingRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms CommonRankingRooms) (CommonRankingRooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(CommonRankingRooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockCommonRankingRoomMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *CommonRankingRoom) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) Find(ctx context.Context, masterRankingId int64, roomId, userId string) (*CommonRankingRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, masterRankingId, roomId, userId)
	ret0, _ := ret[0].(*CommonRankingRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) Find(ctx, masterRankingId, roomId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).Find), ctx, masterRankingId, roomId, userId)
}

// FindByMasterRankingIdAndRoomId mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) FindByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (*CommonRankingRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByMasterRankingIdAndRoomId", ctx, masterRankingId, roomId)
	ret0, _ := ret[0].(*CommonRankingRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByMasterRankingIdAndRoomId indicates an expected call of FindByMasterRankingIdAndRoomId.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) FindByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByMasterRankingIdAndRoomId", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).FindByMasterRankingIdAndRoomId), ctx, masterRankingId, roomId)
}

// FindList mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) FindList(ctx context.Context) (CommonRankingRooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx)
	ret0, _ := ret[0].(CommonRankingRooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) FindList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).FindList), ctx)
}

// FindListByMasterRankingIdAndRoomId mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) FindListByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (CommonRankingRooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindListByMasterRankingIdAndRoomId", ctx, masterRankingId, roomId)
	ret0, _ := ret[0].(CommonRankingRooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindListByMasterRankingIdAndRoomId indicates an expected call of FindListByMasterRankingIdAndRoomId.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) FindListByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindListByMasterRankingIdAndRoomId", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).FindListByMasterRankingIdAndRoomId), ctx, masterRankingId, roomId)
}

// FindOrNil mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) FindOrNil(ctx context.Context, masterRankingId int64, roomId, userId string) (*CommonRankingRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, masterRankingId, roomId, userId)
	ret0, _ := ret[0].(*CommonRankingRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) FindOrNil(ctx, masterRankingId, roomId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).FindOrNil), ctx, masterRankingId, roomId, userId)
}

// FindOrNilByMasterRankingIdAndRoomId mocks base method.
func (m *MockCommonRankingRoomMysqlRepository) FindOrNilByMasterRankingIdAndRoomId(ctx context.Context, masterRankingId int64, roomId string) (*CommonRankingRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNilByMasterRankingIdAndRoomId", ctx, masterRankingId, roomId)
	ret0, _ := ret[0].(*CommonRankingRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNilByMasterRankingIdAndRoomId indicates an expected call of FindOrNilByMasterRankingIdAndRoomId.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) FindOrNilByMasterRankingIdAndRoomId(ctx, masterRankingId, roomId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNilByMasterRankingIdAndRoomId", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).FindOrNilByMasterRankingIdAndRoomId), ctx, masterRankingId, roomId)
}

// Update mocks base method.
func (m_2 *MockCommonRankingRoomMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *CommonRankingRoom) (*CommonRankingRoom, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*CommonRankingRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommonRankingRoomMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommonRankingRoomMysqlRepository)(nil).Update), ctx, tx, m)
}