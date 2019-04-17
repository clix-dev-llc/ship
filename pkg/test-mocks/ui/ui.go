// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mitchellh/cli (interfaces: Ui)

// Package ui is a generated GoMock package.
package ui

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUi is a mock of Ui interface
type MockUi struct {
	ctrl     *gomock.Controller
	recorder *MockUiMockRecorder
}

// MockUiMockRecorder is the mock recorder for MockUi
type MockUiMockRecorder struct {
	mock *MockUi
}

// NewMockUi creates a new mock instance
func NewMockUi(ctrl *gomock.Controller) *MockUi {
	mock := &MockUi{ctrl: ctrl}
	mock.recorder = &MockUiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUi) EXPECT() *MockUiMockRecorder {
	return m.recorder
}

// Ask mocks base method
func (m *MockUi) Ask(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "Ask", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ask indicates an expected call of Ask
func (mr *MockUiMockRecorder) Ask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ask", reflect.TypeOf((*MockUi)(nil).Ask), arg0)
}

// AskSecret mocks base method
func (m *MockUi) AskSecret(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "AskSecret", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AskSecret indicates an expected call of AskSecret
func (mr *MockUiMockRecorder) AskSecret(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AskSecret", reflect.TypeOf((*MockUi)(nil).AskSecret), arg0)
}

// Error mocks base method
func (m *MockUi) Error(arg0 string) {
	m.ctrl.Call(m, "Error", arg0)
}

// Error indicates an expected call of Error
func (mr *MockUiMockRecorder) Error(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockUi)(nil).Error), arg0)
}

// Info mocks base method
func (m *MockUi) Info(arg0 string) {
	m.ctrl.Call(m, "Info", arg0)
}

// Info indicates an expected call of Info
func (mr *MockUiMockRecorder) Info(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockUi)(nil).Info), arg0)
}

// Output mocks base method
func (m *MockUi) Output(arg0 string) {
	m.ctrl.Call(m, "Output", arg0)
}

// Output indicates an expected call of Output
func (mr *MockUiMockRecorder) Output(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Output", reflect.TypeOf((*MockUi)(nil).Output), arg0)
}

// Warn mocks base method
func (m *MockUi) Warn(arg0 string) {
	m.ctrl.Call(m, "Warn", arg0)
}

// Warn indicates an expected call of Warn
func (mr *MockUiMockRecorder) Warn(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockUi)(nil).Warn), arg0)
}
