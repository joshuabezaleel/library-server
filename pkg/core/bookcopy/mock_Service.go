// Code generated by mockery v1.0.0. DO NOT EDIT.

package bookcopy

import mock "github.com/stretchr/testify/mock"

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Create provides a mock function with given fields: bookCopy
func (_m *MockService) Create(bookCopy *BookCopy) (*BookCopy, error) {
	ret := _m.Called(bookCopy)

	var r0 *BookCopy
	if rf, ok := ret.Get(0).(func(*BookCopy) *BookCopy); ok {
		r0 = rf(bookCopy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BookCopy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*BookCopy) error); ok {
		r1 = rf(bookCopy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: bookID
func (_m *MockService) Delete(bookID string) error {
	ret := _m.Called(bookID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: bookID
func (_m *MockService) Get(bookID string) (*BookCopy, error) {
	ret := _m.Called(bookID)

	var r0 *BookCopy
	if rf, ok := ret.Get(0).(func(string) *BookCopy); ok {
		r0 = rf(bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BookCopy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: bookCopy
func (_m *MockService) Update(bookCopy *BookCopy) (*BookCopy, error) {
	ret := _m.Called(bookCopy)

	var r0 *BookCopy
	if rf, ok := ret.Get(0).(func(*BookCopy) *BookCopy); ok {
		r0 = rf(bookCopy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BookCopy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*BookCopy) error); ok {
		r1 = rf(bookCopy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}