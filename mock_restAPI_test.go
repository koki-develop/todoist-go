// Code generated by mockery v2.14.0. DO NOT EDIT.

package todoist

import mock "github.com/stretchr/testify/mock"

// mockRestAPI is an autogenerated mock type for the restAPI type
type mockRestAPI struct {
	mock.Mock
}

// Do provides a mock function with given fields: req
func (_m *mockRestAPI) Do(req *restRequest) (*restResponse, error) {
	ret := _m.Called(req)

	var r0 *restResponse
	if rf, ok := ret.Get(0).(func(*restRequest) *restResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*restResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*restRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockRestAPI interface {
	mock.TestingT
	Cleanup(func())
}

// newMockRestAPI creates a new instance of mockRestAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockRestAPI(t mockConstructorTestingTnewMockRestAPI) *mockRestAPI {
	mock := &mockRestAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
