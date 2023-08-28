// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

// GetInstance provides a mock function with given fields:
func (_m *Router) GetInstance() *echo.Echo {
	ret := _m.Called()

	var r0 *echo.Echo
	if rf, ok := ret.Get(0).(func() *echo.Echo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Echo)
		}
	}

	return r0
}

// NewRouter creates a new instance of Router. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRouter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Router {
	mock := &Router{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
