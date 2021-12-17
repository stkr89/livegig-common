// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import v1 "github.com/stkr89/livegig-common/pb/usersvc/v1"

// UserSvcClient is an autogenerated mock type for the UserSvcClient type
type UserSvcClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *UserSvcClient) Create(ctx context.Context, in *v1.CreateRequest, opts ...grpc.CallOption) (*v1.UserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CreateRequest, ...grpc.CallOption) *v1.UserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, in, opts
func (_m *UserSvcClient) FindByEmail(ctx context.Context, in *v1.FindByEmailRequest, opts ...grpc.CallOption) (*v1.UserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.FindByEmailRequest, ...grpc.CallOption) *v1.UserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.FindByEmailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
