// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/aws-efs-csi-driver/pkg/driver (interfaces: Mounter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mount "k8s.io/kubernetes/pkg/util/mount"
	reflect "reflect"
)

// MockMounter is a mock of Mounter interface
type MockMounter struct {
	ctrl     *gomock.Controller
	recorder *MockMounterMockRecorder
}

// MockMounterMockRecorder is the mock recorder for MockMounter
type MockMounterMockRecorder struct {
	mock *MockMounter
}

// NewMockMounter creates a new mock instance
func NewMockMounter(ctrl *gomock.Controller) *MockMounter {
	mock := &MockMounter{ctrl: ctrl}
	mock.recorder = &MockMounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMounter) EXPECT() *MockMounterMockRecorder {
	return m.recorder
}

// GetMountRefs mocks base method
func (m *MockMounter) GetMountRefs(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMountRefs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMountRefs indicates an expected call of GetMountRefs
func (mr *MockMounterMockRecorder) GetMountRefs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMountRefs", reflect.TypeOf((*MockMounter)(nil).GetMountRefs), arg0)
}

// IsLikelyNotMountPoint mocks base method
func (m *MockMounter) IsLikelyNotMountPoint(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLikelyNotMountPoint", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLikelyNotMountPoint indicates an expected call of IsLikelyNotMountPoint
func (mr *MockMounterMockRecorder) IsLikelyNotMountPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLikelyNotMountPoint", reflect.TypeOf((*MockMounter)(nil).IsLikelyNotMountPoint), arg0)
}

// List mocks base method
func (m *MockMounter) List() ([]mount.MountPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]mount.MountPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockMounterMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMounter)(nil).List))
}

// MakeDir mocks base method
func (m *MockMounter) MakeDir(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeDir", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeDir indicates an expected call of MakeDir
func (mr *MockMounterMockRecorder) MakeDir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeDir", reflect.TypeOf((*MockMounter)(nil).MakeDir), arg0)
}

// Mount mocks base method
func (m *MockMounter) Mount(arg0, arg1, arg2 string, arg3 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mount indicates an expected call of Mount
func (mr *MockMounterMockRecorder) Mount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockMounter)(nil).Mount), arg0, arg1, arg2, arg3)
}

// Unmount mocks base method
func (m *MockMounter) Unmount(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmount indicates an expected call of Unmount
func (mr *MockMounterMockRecorder) Unmount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockMounter)(nil).Unmount), arg0)
}
