// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Position is an autogenerated mock type for the Position type
type Position struct {
	mock.Mock
}

// GetColumn provides a mock function with given fields:
func (_m *Position) GetColumn() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetLine provides a mock function with given fields:
func (_m *Position) GetLine() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
