// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kube-startup-cpu-boost/internal/boost (interfaces: Manager)
//
// Generated by this command:
//
//	mockgen -package mock --copyright_file hack/boilerplate.go.txt --destination internal/mock/boost_manager.go github.com/google/kube-startup-cpu-boost/internal/boost Manager
//

package mock

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/google/kube-startup-cpu-boost/api/v1alpha1"
	boost "github.com/google/kube-startup-cpu-boost/internal/boost"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
	isgomock struct{}
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddRegularCPUBoost mocks base method.
func (m *MockManager) AddRegularCPUBoost(ctx context.Context, boost boost.StartupCPUBoost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRegularCPUBoost", ctx, boost)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRegularCPUBoost indicates an expected call of AddRegularCPUBoost.
func (mr *MockManagerMockRecorder) AddRegularCPUBoost(ctx, boost any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRegularCPUBoost", reflect.TypeOf((*MockManager)(nil).AddRegularCPUBoost), ctx, boost)
}

// DeletePod mocks base method.
func (m *MockManager) DeletePod(ctx context.Context, pod *v1.Pod) (boost.StartupCPUBoost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", ctx, pod)
	ret0, _ := ret[0].(boost.StartupCPUBoost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePod indicates an expected call of DeletePod.
func (mr *MockManagerMockRecorder) DeletePod(ctx, pod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockManager)(nil).DeletePod), ctx, pod)
}

// DeleteRegularCPUBoost mocks base method.
func (m *MockManager) DeleteRegularCPUBoost(ctx context.Context, name, namespace string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteRegularCPUBoost", ctx, name, namespace)
}

// DeleteRegularCPUBoost indicates an expected call of DeleteRegularCPUBoost.
func (mr *MockManagerMockRecorder) DeleteRegularCPUBoost(ctx, name, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRegularCPUBoost", reflect.TypeOf((*MockManager)(nil).DeleteRegularCPUBoost), ctx, name, namespace)
}

// GetCPUBoostForPod mocks base method.
func (m *MockManager) GetCPUBoostForPod(ctx context.Context, pod *v1.Pod) (boost.StartupCPUBoost, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCPUBoostForPod", ctx, pod)
	ret0, _ := ret[0].(boost.StartupCPUBoost)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCPUBoostForPod indicates an expected call of GetCPUBoostForPod.
func (mr *MockManagerMockRecorder) GetCPUBoostForPod(ctx, pod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCPUBoostForPod", reflect.TypeOf((*MockManager)(nil).GetCPUBoostForPod), ctx, pod)
}

// GetRegularCPUBoost mocks base method.
func (m *MockManager) GetRegularCPUBoost(ctx context.Context, name, namespace string) (boost.StartupCPUBoost, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRegularCPUBoost", ctx, name, namespace)
	ret0, _ := ret[0].(boost.StartupCPUBoost)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetRegularCPUBoost indicates an expected call of GetRegularCPUBoost.
func (mr *MockManagerMockRecorder) GetRegularCPUBoost(ctx, name, namespace any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegularCPUBoost", reflect.TypeOf((*MockManager)(nil).GetRegularCPUBoost), ctx, name, namespace)
}

// IsRunning mocks base method.
func (m *MockManager) IsRunning(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning.
func (mr *MockManagerMockRecorder) IsRunning(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockManager)(nil).IsRunning), ctx)
}

// SetStartupCPUBoostReconciler mocks base method.
func (m *MockManager) SetStartupCPUBoostReconciler(reconciler reconcile.TypedReconciler[reconcile.Request]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStartupCPUBoostReconciler", reconciler)
}

// SetStartupCPUBoostReconciler indicates an expected call of SetStartupCPUBoostReconciler.
func (mr *MockManagerMockRecorder) SetStartupCPUBoostReconciler(reconciler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStartupCPUBoostReconciler", reflect.TypeOf((*MockManager)(nil).SetStartupCPUBoostReconciler), reconciler)
}

// Start mocks base method.
func (m *MockManager) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockManagerMockRecorder) Start(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockManager)(nil).Start), ctx)
}

// UpdateRegularCPUBoost mocks base method.
func (m *MockManager) UpdateRegularCPUBoost(ctx context.Context, spec *v1alpha1.StartupCPUBoost) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRegularCPUBoost", ctx, spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRegularCPUBoost indicates an expected call of UpdateRegularCPUBoost.
func (mr *MockManagerMockRecorder) UpdateRegularCPUBoost(ctx, spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRegularCPUBoost", reflect.TypeOf((*MockManager)(nil).UpdateRegularCPUBoost), ctx, spec)
}

// UpsertPod mocks base method.
func (m *MockManager) UpsertPod(ctx context.Context, pod *v1.Pod) (boost.StartupCPUBoost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPod", ctx, pod)
	ret0, _ := ret[0].(boost.StartupCPUBoost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertPod indicates an expected call of UpsertPod.
func (mr *MockManagerMockRecorder) UpsertPod(ctx, pod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPod", reflect.TypeOf((*MockManager)(nil).UpsertPod), ctx, pod)
}
