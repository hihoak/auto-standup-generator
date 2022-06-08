// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package jirer is a generated GoMock package.
package jirer

import (
	reflect "reflect"

	jira "github.com/andygrunwald/go-jira"
	gomock "github.com/golang/mock/gomock"
)

// MockJirer is a mock of Jirer interface.
type MockJirer struct {
	ctrl     *gomock.Controller
	recorder *MockJirerMockRecorder
}

// MockJirerMockRecorder is the mock recorder for MockJirer.
type MockJirerMockRecorder struct {
	mock *MockJirer
}

// NewMockJirer creates a new mock instance.
func NewMockJirer(ctrl *gomock.Controller) *MockJirer {
	mock := &MockJirer{ctrl: ctrl}
	mock.recorder = &MockJirerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJirer) EXPECT() *MockJirerMockRecorder {
	return m.recorder
}

// GetIssue mocks base method.
func (m *MockJirer) GetIssue(key string, options *jira.GetQueryOptions) (*jira.Issue, *jira.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssue", key, options)
	ret0, _ := ret[0].(*jira.Issue)
	ret1, _ := ret[1].(*jira.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIssue indicates an expected call of GetIssue.
func (mr *MockJirerMockRecorder) GetIssue(key, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssue", reflect.TypeOf((*MockJirer)(nil).GetIssue), key, options)
}

// SearchIssue mocks base method.
func (m *MockJirer) SearchIssue(jql string, options *jira.SearchOptions) ([]jira.Issue, *jira.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIssue", jql, options)
	ret0, _ := ret[0].([]jira.Issue)
	ret1, _ := ret[1].(*jira.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchIssue indicates an expected call of SearchIssue.
func (mr *MockJirerMockRecorder) SearchIssue(jql, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIssue", reflect.TypeOf((*MockJirer)(nil).SearchIssue), jql, options)
}
