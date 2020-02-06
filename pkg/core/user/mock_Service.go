// Code generated by mockery v1.0.0. DO NOT EDIT.

package user

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddFine provides a mock function with given fields: userID, fine
func (_m *MockService) AddFine(userID string, fine uint32) (uint32, error) {
	ret := _m.Called(userID, fine)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(string, uint32) uint32); ok {
		r0 = rf(userID, fine)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint32) error); ok {
		r1 = rf(userID, fine)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: user
func (_m *MockService) Create(user *User) (*User, error) {
	ret := _m.Called(user)

	var r0 *User
	if rf, ok := ret.Get(0).(func(*User) *User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID
func (_m *MockService) Delete(userID string) error {
	ret := _m.Called(userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: userID
func (_m *MockService) Get(userID string) (*User, error) {
	ret := _m.Called(userID)

	var r0 *User
	if rf, ok := ret.Get(0).(func(string) *User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRole provides a mock function with given fields: username
func (_m *MockService) GetRole(username string) (string, error) {
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

// GetTotalFine provides a mock function with given fields: userID
func (_m *MockService) GetTotalFine(userID string) (uint32, error) {
	ret := _m.Called(userID)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserIDByUsername provides a mock function with given fields: username
func (_m *MockService) GetUserIDByUsername(username string) (string, error) {
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

// Update provides a mock function with given fields: user
func (_m *MockService) Update(user *User) (*User, error) {
	ret := _m.Called(user)

	var r0 *User
	if rf, ok := ret.Get(0).(func(*User) *User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
