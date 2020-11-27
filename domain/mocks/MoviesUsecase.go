// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinbyte/movies/domain"
)

// MoviesUsecase is an autogenerated mock type for the MoviesUsecase type
type MoviesUsecase struct {
	mock.Mock
}

// DetailProcess provides a mock function with given fields: ctx, param
func (_m *MoviesUsecase) DetailProcess(ctx context.Context, param domain.DetailParam) domain.BaseResponse {
	ret := _m.Called(ctx, param)

	var r0 domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, domain.DetailParam) domain.BaseResponse); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(domain.BaseResponse)
	}

	return r0
}

// SearchProcess provides a mock function with given fields: ctx, param
func (_m *MoviesUsecase) SearchProcess(ctx context.Context, param domain.SearchParam) domain.BaseResponse {
	ret := _m.Called(ctx, param)

	var r0 domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, domain.SearchParam) domain.BaseResponse); ok {
		r0 = rf(ctx, param)
	} else {
		r0 = ret.Get(0).(domain.BaseResponse)
	}

	return r0
}
