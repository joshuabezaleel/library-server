// Code generated by mockery v1.0.0. DO NOT EDIT.

package borrowing

import mock "github.com/stretchr/testify/mock"

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// Borrow provides a mock function with given fields: borrow
func (_m *MockRepository) Borrow(borrow *Borrow) (*Borrow, error) {
	ret := _m.Called(borrow)

	var r0 *Borrow
	if rf, ok := ret.Get(0).(func(*Borrow) *Borrow); ok {
		r0 = rf(borrow)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Borrow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Borrow) error); ok {
		r1 = rf(borrow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckBorrowed provides a mock function with given fields: bookCopyID
func (_m *MockRepository) CheckBorrowed(bookCopyID string) (bool, error) {
	ret := _m.Called(bookCopyID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(bookCopyID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(bookCopyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: borrowID
func (_m *MockRepository) Get(borrowID string) (*Borrow, error) {
	ret := _m.Called(borrowID)

	var r0 *Borrow
	if rf, ok := ret.Get(0).(func(string) *Borrow); ok {
		r0 = rf(borrowID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Borrow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(borrowID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserIDAndBookCopyID provides a mock function with given fields: userID, bookCopyID
func (_m *MockRepository) GetByUserIDAndBookCopyID(userID string, bookCopyID string) (*Borrow, error) {
	ret := _m.Called(userID, bookCopyID)

	var r0 *Borrow
	if rf, ok := ret.Get(0).(func(string, string) *Borrow); ok {
		r0 = rf(userID, bookCopyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Borrow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, bookCopyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Return provides a mock function with given fields: borrow
func (_m *MockRepository) Return(borrow *Borrow) (*Borrow, error) {
	ret := _m.Called(borrow)

	var r0 *Borrow
	if rf, ok := ret.Get(0).(func(*Borrow) *Borrow); ok {
		r0 = rf(borrow)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Borrow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Borrow) error); ok {
		r1 = rf(borrow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
