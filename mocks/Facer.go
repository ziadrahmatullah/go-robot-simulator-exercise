// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	event "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/go-robot-simulator-exercise/event"
	mock "github.com/stretchr/testify/mock"
)

// Facer is an autogenerated mock type for the Facer type
type Facer struct {
	mock.Mock
}

// face provides a mock function with given fields: _a0
func (_m *Facer) face(_a0 *event.Facing) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewFacer interface {
	mock.TestingT
	Cleanup(func())
}

// NewFacer creates a new instance of Facer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFacer(t mockConstructorTestingTNewFacer) *Facer {
	mock := &Facer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
