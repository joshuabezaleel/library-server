// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import user "github.com/joshuabezaleel/library-server/pkg/core/user"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddFine provides a mock function with given fields: userID, fine
func (_m *Repository) AddFine(userID string, fine uint32) error {
	ret := _m.Called(userID, fine)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32) error); ok {
		r0 = rf(userID, fine)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckLibrarian provides a mock function with given fields: userID
func (_m *Repository) CheckLibrarian(userID string) (string, error) {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID
func (_m *Repository) Delete(userID string) error {
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
func (_m *Repository) Get(userID string) (*user.User, error) {
	ret := _m.Called(userID)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string) *user.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
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

// GetIDByUsername provides a mock function with given fields: username
func (_m *Repository) GetIDByUsername(username string) (string, error) {
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
func (_m *Repository) GetTotalFine(userID string) (uint32, error) {
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

// Save provides a mock function with given fields: _a0
func (_m *Repository) Save(_a0 *user.User) (*user.User, error) {
	ret := _m.Called(_a0)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(*user.User) *user.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *Repository) Update(_a0 *user.User) (*user.User, error) {
	ret := _m.Called(_a0)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(*user.User) *user.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*user.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
