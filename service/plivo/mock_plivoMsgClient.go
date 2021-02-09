// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package plivo

import (
	plivo "github.com/plivo/plivo-go"
	mock "github.com/stretchr/testify/mock"
)

// mockPlivoMsgClient is an autogenerated mock type for the plivoMsgClient type
type mockPlivoMsgClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *mockPlivoMsgClient) Create(_a0 plivo.MessageCreateParams) (*plivo.MessageCreateResponseBody, error) {
	ret := _m.Called(_a0)

	var r0 *plivo.MessageCreateResponseBody
	if rf, ok := ret.Get(0).(func(plivo.MessageCreateParams) *plivo.MessageCreateResponseBody); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*plivo.MessageCreateResponseBody)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(plivo.MessageCreateParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
