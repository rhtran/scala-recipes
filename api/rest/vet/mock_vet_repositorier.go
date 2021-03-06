// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package vet

import mock "github.com/stretchr/testify/mock"

// MockVetRepositorier is an autogenerated mock type for the VetRepositorier type
type MockVetRepositorier struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *MockVetRepositorier) FindAll() ([]Vet, error) {
	ret := _m.Called()

	var r0 []Vet
	if rf, ok := ret.Get(0).(func() []Vet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Vet)
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

// FindAllPreload provides a mock function with given fields:
func (_m *MockVetRepositorier) FindAllPreload() ([]Vet, error) {
	ret := _m.Called()

	var r0 []Vet
	if rf, ok := ret.Get(0).(func() []Vet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Vet)
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

// FindById provides a mock function with given fields: id
func (_m *MockVetRepositorier) FindById(id int) (*Vet, error) {
	ret := _m.Called(id)

	var r0 *Vet
	if rf, ok := ret.Get(0).(func(int) *Vet); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Vet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByLastName provides a mock function with given fields: lastName
func (_m *MockVetRepositorier) FindByLastName(lastName string) ([]Vet, error) {
	ret := _m.Called(lastName)

	var r0 []Vet
	if rf, ok := ret.Get(0).(func(string) []Vet); ok {
		r0 = rf(lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Vet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: vet
func (_m *MockVetRepositorier) Insert(vet *Vet) (*Vet, error) {
	ret := _m.Called(vet)

	var r0 *Vet
	if rf, ok := ret.Get(0).(func(*Vet) *Vet); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Vet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: vet
func (_m *MockVetRepositorier) Update(vet *Vet) (*Vet, error) {
	ret := _m.Called(vet)

	var r0 *Vet
	if rf, ok := ret.Get(0).(func(*Vet) *Vet); ok {
		r0 = rf(vet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Vet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Vet) error); ok {
		r1 = rf(vet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
