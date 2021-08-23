// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/kube-scheduler-simulator/node (interfaces: PodService)

// Package mock_node is a generated GoMock package.
package mock_node

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockPodService is a mock of PodService interface.
type MockPodService struct {
	ctrl     *gomock.Controller
	recorder *MockPodServiceMockRecorder
}

// MockPodServiceMockRecorder is the mock recorder for MockPodService.
type MockPodServiceMockRecorder struct {
	mock *MockPodService
}

// NewMockPodService creates a new mock instance.
func NewMockPodService(ctrl *gomock.Controller) *MockPodService {
	mock := &MockPodService{ctrl: ctrl}
	mock.recorder = &MockPodServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodService) EXPECT() *MockPodServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPodService) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPodServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPodService)(nil).Delete), arg0, arg1)
}

// List mocks base method.
func (m *MockPodService) List(arg0 context.Context) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPodServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPodService)(nil).List), arg0)
}
