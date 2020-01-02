// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import generated "github.com/pPrecel/BeerKongServer/pkg/graphql/generated"
import mock "github.com/stretchr/testify/mock"

// ResolverRoot is an autogenerated mock type for the ResolverRoot type
type ResolverRoot struct {
	mock.Mock
}

// League provides a mock function with given fields:
func (_m *ResolverRoot) League() generated.LeagueResolver {
	ret := _m.Called()

	var r0 generated.LeagueResolver
	if rf, ok := ret.Get(0).(func() generated.LeagueResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.LeagueResolver)
		}
	}

	return r0
}

// Match provides a mock function with given fields:
func (_m *ResolverRoot) Match() generated.MatchResolver {
	ret := _m.Called()

	var r0 generated.MatchResolver
	if rf, ok := ret.Get(0).(func() generated.MatchResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.MatchResolver)
		}
	}

	return r0
}

// Mutation provides a mock function with given fields:
func (_m *ResolverRoot) Mutation() generated.MutationResolver {
	ret := _m.Called()

	var r0 generated.MutationResolver
	if rf, ok := ret.Get(0).(func() generated.MutationResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.MutationResolver)
		}
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *ResolverRoot) Query() generated.QueryResolver {
	ret := _m.Called()

	var r0 generated.QueryResolver
	if rf, ok := ret.Get(0).(func() generated.QueryResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.QueryResolver)
		}
	}

	return r0
}

// Team provides a mock function with given fields:
func (_m *ResolverRoot) Team() generated.TeamResolver {
	ret := _m.Called()

	var r0 generated.TeamResolver
	if rf, ok := ret.Get(0).(func() generated.TeamResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.TeamResolver)
		}
	}

	return r0
}

// User provides a mock function with given fields:
func (_m *ResolverRoot) User() generated.UserResolver {
	ret := _m.Called()

	var r0 generated.UserResolver
	if rf, ok := ret.Get(0).(func() generated.UserResolver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.UserResolver)
		}
	}

	return r0
}
