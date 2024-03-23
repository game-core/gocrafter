// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_item_box_repository.gen.go

// Package userItemBox is a generated GoMock package.
package userItemBox

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserItemBoxMysqlRepository is a mock of UserItemBoxMysqlRepository interface.
type MockUserItemBoxMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserItemBoxMysqlRepositoryMockRecorder
}

// MockUserItemBoxMysqlRepositoryMockRecorder is the mock recorder for MockUserItemBoxMysqlRepository.
type MockUserItemBoxMysqlRepositoryMockRecorder struct {
	mock *MockUserItemBoxMysqlRepository
}

// NewMockUserItemBoxMysqlRepository creates a new mock instance.
func NewMockUserItemBoxMysqlRepository(ctrl *gomock.Controller) *MockUserItemBoxMysqlRepository {
	mock := &MockUserItemBoxMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserItemBoxMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserItemBoxMysqlRepository) EXPECT() *MockUserItemBoxMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserItemBoxMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *UserItemBox) (*UserItemBox, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*UserItemBox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockUserItemBoxMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms UserItemBoxes) (UserItemBoxes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserItemBoxes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockUserItemBoxMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *UserItemBox) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockUserItemBoxMysqlRepository) Find(ctx context.Context, userId string, masterItemId int64) (*UserItemBox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, userId, masterItemId)
	ret0, _ := ret[0].(*UserItemBox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) Find(ctx, userId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).Find), ctx, userId, masterItemId)
}

// FindList mocks base method.
func (m *MockUserItemBoxMysqlRepository) FindList(ctx context.Context, userId string) (UserItemBoxes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx, userId)
	ret0, _ := ret[0].(UserItemBoxes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) FindList(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).FindList), ctx, userId)
}

// FindOrNil mocks base method.
func (m *MockUserItemBoxMysqlRepository) FindOrNil(ctx context.Context, userId string, masterItemId int64) (*UserItemBox, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, userId, masterItemId)
	ret0, _ := ret[0].(*UserItemBox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) FindOrNil(ctx, userId, masterItemId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).FindOrNil), ctx, userId, masterItemId)
}

// Update mocks base method.
func (m_2 *MockUserItemBoxMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *UserItemBox) (*UserItemBox, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*UserItemBox)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserItemBoxMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserItemBoxMysqlRepository)(nil).Update), ctx, tx, m)
}
