// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package visit

import mock "github.com/stretchr/testify/mock"

// MockVisitServicer is an autogenerated mock type for the VisitServicer type
type MockVisitServicer struct {
	mock.Mock
}

// GetAllVisits provides a mock function with given fields:
func (_m *MockVisitServicer) GetAllVisits() ([]VisitResponse, error) {
	ret := _m.Called()

	var r0 []VisitResponse
	if rf, ok := ret.Get(0).(func() []VisitResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]VisitResponse)
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

// GetVisitById provides a mock function with given fields: id
func (_m *MockVisitServicer) GetVisitById(id int) (*VisitResponse, error) {
	ret := _m.Called(id)

	var r0 *VisitResponse
	if rf, ok := ret.Get(0).(func(int) *VisitResponse); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*VisitResponse)
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
