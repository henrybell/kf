// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/GoogleCloudPlatform/kf/pkg/kf/buildpacks/fake (interfaces: BuildpackLister)

// Package fake is a generated GoMock package.
package fake

import (
	buildpacks "github.com/GoogleCloudPlatform/kf/pkg/kf/buildpacks"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// FakeBuildpackLister is a mock of BuildpackLister interface
type FakeBuildpackLister struct {
	ctrl     *gomock.Controller
	recorder *FakeBuildpackListerMockRecorder
}

// FakeBuildpackListerMockRecorder is the mock recorder for FakeBuildpackLister
type FakeBuildpackListerMockRecorder struct {
	mock *FakeBuildpackLister
}

// NewFakeBuildpackLister creates a new mock instance
func NewFakeBuildpackLister(ctrl *gomock.Controller) *FakeBuildpackLister {
	mock := &FakeBuildpackLister{ctrl: ctrl}
	mock.recorder = &FakeBuildpackListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeBuildpackLister) EXPECT() *FakeBuildpackListerMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *FakeBuildpackLister) List(arg0 ...buildpacks.BuildpackListOption) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *FakeBuildpackListerMockRecorder) List(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeBuildpackLister)(nil).List), arg0...)
}