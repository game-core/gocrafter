// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_transaction_redis_repository.gen.go

// Package userTransaction is a generated GoMock package.
package userTransaction

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	redis "github.com/redis/go-redis/v9"
)

// MockUserTransactionRedisRepository is a mock of UserTransactionRedisRepository interface.
type MockUserTransactionRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserTransactionRedisRepositoryMockRecorder
}

// MockUserTransactionRedisRepositoryMockRecorder is the mock recorder for MockUserTransactionRedisRepository.
type MockUserTransactionRedisRepositoryMockRecorder struct {
	mock *MockUserTransactionRedisRepository
}

// NewMockUserTransactionRedisRepository creates a new mock instance.
func NewMockUserTransactionRedisRepository(ctrl *gomock.Controller) *MockUserTransactionRedisRepository {
	mock := &MockUserTransactionRedisRepository{ctrl: ctrl}
	mock.recorder = &MockUserTransactionRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserTransactionRedisRepository) EXPECT() *MockUserTransactionRedisRepositoryMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockUserTransactionRedisRepository) Begin() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// Begin indicates an expected call of Begin.
func (mr *MockUserTransactionRedisRepositoryMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockUserTransactionRedisRepository)(nil).Begin))
}

// Commit mocks base method.
func (m *MockUserTransactionRedisRepository) Commit(ctx context.Context, tx redis.Pipeliner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockUserTransactionRedisRepositoryMockRecorder) Commit(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockUserTransactionRedisRepository)(nil).Commit), ctx, tx)
}

// Rollback mocks base method.
func (m *MockUserTransactionRedisRepository) Rollback(tx redis.Pipeliner) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rollback", tx)
}

// Rollback indicates an expected call of Rollback.
func (mr *MockUserTransactionRedisRepositoryMockRecorder) Rollback(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockUserTransactionRedisRepository)(nil).Rollback), tx)
}