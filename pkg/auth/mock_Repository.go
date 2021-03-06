// Code generated by mockery v1.0.0. DO NOT EDIT.

package auth

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// GetPassword provides a mock function with given fields: username
func (_m *MockRepository) GetPassword(username string) (string, error) {
	ret := _m.Called(username)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
