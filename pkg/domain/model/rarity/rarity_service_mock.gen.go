// Code generated by MockGen. DO NOT EDIT.
// Source: ./rarity_service.go

// Package rarity is a generated GoMock package.
package rarity

import (
	gomock "github.com/golang/mock/gomock"
)

// MockRarityService is a mock of RarityService interface.
type MockRarityService struct {
	ctrl     *gomock.Controller
	recorder *MockRarityServiceMockRecorder
}

// MockRarityServiceMockRecorder is the mock recorder for MockRarityService.
type MockRarityServiceMockRecorder struct {
	mock *MockRarityService
}

// NewMockRarityService creates a new mock instance.
func NewMockRarityService(ctrl *gomock.Controller) *MockRarityService {
	mock := &MockRarityService{ctrl: ctrl}
	mock.recorder = &MockRarityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRarityService) EXPECT() *MockRarityServiceMockRecorder {
	return m.recorder
}