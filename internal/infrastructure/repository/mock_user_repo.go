package repository

import (
	"api/internal/entity"

	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (_m *MockUserRepo) GetByUUID(ctx context.Context, uid string) (*entity.User, error) {

	ret := _m.Called(ctx, uid)

	var r0 *entity.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*entity.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (_m *MockUserRepo) Update(ctx context.Context, u string, na string) error {
	ret := _m.Called(ctx, u, na)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, u, na)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockUserRepo) GetAll(ctx context.Context) ([]*entity.User, error) {
	ret := _m.Called(ctx)

	var r0 []*entity.User
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(2)
	}

	return r0, r1
}

func (_m *MockUserRepo) Set(ctx context.Context, u *entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Error(0)
		}
	}

	return nil, r0
}

func (_m *MockUserRepo) MakeFriends(ctx context.Context, suid string, tuid string) error {
	ret := _m.Called(ctx, suid, tuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, suid, tuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockUserRepo) GetAllFriends(ctx context.Context, uid string) ([]*entity.Friends, error) {
	ret := _m.Called(ctx, uid)

	var r0 []*entity.Friends
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Friends); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Friends)
		}
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(2)
	}

	return r0, r1
}

func (_m *MockUserRepo) Delete(ctx context.Context, uid string) error {
	ret := _m.Called(ctx, uid)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, uid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
