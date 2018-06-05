package mocks

import common "github.com/opencontrol/compliance-masonry/pkg/lib/common"
import mock "github.com/stretchr/testify/mock"

// Section is an autogenerated mock type for the Section type
type Section struct {
	mock.Mock
}

// GetKey provides a mock function with given fields:
func (_m *Section) GetKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetText provides a mock function with given fields:
func (_m *Section) GetText() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

var _ common.Section = (*Section)(nil)
