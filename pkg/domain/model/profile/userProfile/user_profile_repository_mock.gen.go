// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_profile_repository.gen.go

// Package userProfile is a generated GoMock package.
package userProfile

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockUserProfileMysqlRepository is a mock of UserProfileMysqlRepository interface.
type MockUserProfileMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserProfileMysqlRepositoryMockRecorder
}

// MockUserProfileMysqlRepositoryMockRecorder is the mock recorder for MockUserProfileMysqlRepository.
type MockUserProfileMysqlRepositoryMockRecorder struct {
	mock *MockUserProfileMysqlRepository
}

// NewMockUserProfileMysqlRepository creates a new mock instance.
func NewMockUserProfileMysqlRepository(ctrl *gomock.Controller) *MockUserProfileMysqlRepository {
	mock := &MockUserProfileMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserProfileMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserProfileMysqlRepository) EXPECT() *MockUserProfileMysqlRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockUserProfileMysqlRepository) Create(ctx context.Context, tx *gorm.DB, m *UserProfile) (*UserProfile, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, tx, m)
	ret0, _ := ret[0].(*UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) Create(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).Create), ctx, tx, m)
}

// CreateList mocks base method.
func (m *MockUserProfileMysqlRepository) CreateList(ctx context.Context, tx *gorm.DB, ms UserProfiles) (UserProfiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateList", ctx, tx, ms)
	ret0, _ := ret[0].(UserProfiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateList indicates an expected call of CreateList.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) CreateList(ctx, tx, ms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateList", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).CreateList), ctx, tx, ms)
}

// Delete mocks base method.
func (m_2 *MockUserProfileMysqlRepository) Delete(ctx context.Context, tx *gorm.DB, m *UserProfile) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Delete", ctx, tx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) Delete(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).Delete), ctx, tx, m)
}

// Find mocks base method.
func (m *MockUserProfileMysqlRepository) Find(ctx context.Context, userId string) (*UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, userId)
	ret0, _ := ret[0].(*UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) Find(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).Find), ctx, userId)
}

// FindList mocks base method.
func (m *MockUserProfileMysqlRepository) FindList(ctx context.Context, userId string) (UserProfiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindList", ctx, userId)
	ret0, _ := ret[0].(UserProfiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindList indicates an expected call of FindList.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) FindList(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindList", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).FindList), ctx, userId)
}

// FindOrNil mocks base method.
func (m *MockUserProfileMysqlRepository) FindOrNil(ctx context.Context, userId string) (*UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrNil", ctx, userId)
	ret0, _ := ret[0].(*UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrNil indicates an expected call of FindOrNil.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) FindOrNil(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrNil", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).FindOrNil), ctx, userId)
}

// Update mocks base method.
func (m_2 *MockUserProfileMysqlRepository) Update(ctx context.Context, tx *gorm.DB, m *UserProfile) (*UserProfile, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, tx, m)
	ret0, _ := ret[0].(*UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserProfileMysqlRepositoryMockRecorder) Update(ctx, tx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserProfileMysqlRepository)(nil).Update), ctx, tx, m)
}
