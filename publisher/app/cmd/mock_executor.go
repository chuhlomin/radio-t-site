// Code generated by mockery v1.0.0. DO NOT EDIT.

package cmd

import mock "github.com/stretchr/testify/mock"

// MockExecutor is an autogenerated mock type for the Executor type
type MockExecutor struct {
	mock.Mock
}

// Run provides a mock function with given fields: cmd, params
func (_m *MockExecutor) Run(cmd string, params ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, cmd)
	_ca = append(_ca, params...)
	_m.Called(_ca...)
}
