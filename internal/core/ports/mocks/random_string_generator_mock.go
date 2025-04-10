// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/random_string_generator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRandomStringGenerator is a mock of IRandomStringGenerator interface.
type MockIRandomStringGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockIRandomStringGeneratorMockRecorder
}

// MockIRandomStringGeneratorMockRecorder is the mock recorder for MockIRandomStringGenerator.
type MockIRandomStringGeneratorMockRecorder struct {
	mock *MockIRandomStringGenerator
}

// NewMockIRandomStringGenerator creates a new mock instance.
func NewMockIRandomStringGenerator(ctrl *gomock.Controller) *MockIRandomStringGenerator {
	mock := &MockIRandomStringGenerator{ctrl: ctrl}
	mock.recorder = &MockIRandomStringGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRandomStringGenerator) EXPECT() *MockIRandomStringGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockIRandomStringGenerator) Generate(length int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", length)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockIRandomStringGeneratorMockRecorder) Generate(length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockIRandomStringGenerator)(nil).Generate), length)
}
