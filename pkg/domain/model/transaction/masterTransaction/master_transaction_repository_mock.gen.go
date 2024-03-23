// Code generated by MockGen. DO NOT EDIT.
// Source: ./master_transaction_repository.gen.go

// Package masterTransaction is a generated GoMock package.
package masterTransaction

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockMasterTransactionMysqlRepository is a mock of MasterTransactionMysqlRepository interface.
type MockMasterTransactionMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMasterTransactionMysqlRepositoryMockRecorder
}

// MockMasterTransactionMysqlRepositoryMockRecorder is the mock recorder for MockMasterTransactionMysqlRepository.
type MockMasterTransactionMysqlRepositoryMockRecorder struct {
	mock *MockMasterTransactionMysqlRepository
}

// NewMockMasterTransactionMysqlRepository creates a new mock instance.
func NewMockMasterTransactionMysqlRepository(ctrl *gomock.Controller) *MockMasterTransactionMysqlRepository {
	mock := &MockMasterTransactionMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockMasterTransactionMysqlRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMasterTransactionMysqlRepository) EXPECT() *MockMasterTransactionMysqlRepositoryMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockMasterTransactionMysqlRepository) Begin(ctx context.Context) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", ctx)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockMasterTransactionMysqlRepositoryMockRecorder) Begin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockMasterTransactionMysqlRepository)(nil).Begin), ctx)
}

// Commit mocks base method.
func (m *MockMasterTransactionMysqlRepository) Commit(ctx context.Context, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockMasterTransactionMysqlRepositoryMockRecorder) Commit(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockMasterTransactionMysqlRepository)(nil).Commit), ctx, tx)
}

// Rollback mocks base method.
func (m *MockMasterTransactionMysqlRepository) Rollback(ctx context.Context, tx *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockMasterTransactionMysqlRepositoryMockRecorder) Rollback(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockMasterTransactionMysqlRepository)(nil).Rollback), ctx, tx)
}
