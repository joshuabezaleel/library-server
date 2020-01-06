// Code generated by mockery v1.0.0. DO NOT EDIT.

package user

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// AddFine provides a mock function with given fields: userID, fine
func (_m *MockRepository) AddFine(userID string, fine uint32) error {
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
func (_m *MockRepository) CheckLibrarian(userID string) (string, error) {
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
func (_m *MockRepository) Delete(userID string) error {
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
func (_m *MockRepository) Get(userID string) (*User, error) {
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

// GetIDByUsername provides a mock function with given fields: username
func (_m *MockRepository) GetIDByUsername(username string) (string, error) {
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
func (_m *MockRepository) GetTotalFine(userID string) (uint32, error) {
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

// Save provides a mock function with given fields: user
func (_m *MockRepository) Save(user *User) (*User, error) {
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

// Update provides a mock function with given fields: user
func (_m *MockRepository) Update(user *User) (*User, error) {
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