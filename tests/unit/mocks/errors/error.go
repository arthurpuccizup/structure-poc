// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Error is an autogenerated mock type for the Error type
type Error struct {
	mock.Mock
}

// AddMeta provides a mock function with given fields: key, value
func (_m *Error) AddMeta(key string, value string) *errors.sensitiveError {
	ret := _m.Called(key, value)

	var r0 *errors.sensitiveError
	if rf, ok := ret.Get(0).(func(string, string) *errors.sensitiveError); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.sensitiveError)
		}
	}

	return r0
}

// AddOperation provides a mock function with given fields: operation
func (_m *Error) AddOperation(operation string) *errors.sensitiveError {
	ret := _m.Called(operation)

	var r0 *errors.sensitiveError
	if rf, ok := ret.Get(0).(func(string) *errors.sensitiveError); ok {
		r0 = rf(operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errors.sensitiveError)
		}
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *Error) Error() errors.sensitiveError {
	ret := _m.Called()

	var r0 errors.sensitiveError
	if rf, ok := ret.Get(0).(func() errors.sensitiveError); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(errors.sensitiveError)
	}

	return r0
}

// Marshal provides a mock function with given fields:
func (_m *Error) Marshal() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SensitiveError provides a mock function with given fields:
func (_m *Error) SensitiveError() errors.customError {
	ret := _m.Called()

	var r0 errors.customError
	if rf, ok := ret.Get(0).(func() errors.customError); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(errors.customError)
	}

	return r0
}