// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"
)

// BlockHeightWatcher is an autogenerated mock type for the BlockHeightWatcher type
type BlockHeightWatcher struct {
	mock.Mock
}

// GetMetrics provides a mock function with given fields: labels
func (_m *BlockHeightWatcher) GetMetrics(labels map[string]string) []prometheus.Collector {
	ret := _m.Called(labels)

	var r0 []prometheus.Collector
	if rf, ok := ret.Get(0).(func(map[string]string) []prometheus.Collector); ok {
		r0 = rf(labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]prometheus.Collector)
		}
	}

	return r0
}

// Subscribe provides a mock function with given fields:
func (_m *BlockHeightWatcher) Subscribe() <-chan uint64 {
	ret := _m.Called()

	var r0 <-chan uint64
	if rf, ok := ret.Get(0).(func() <-chan uint64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan uint64)
		}
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: ch
func (_m *BlockHeightWatcher) Unsubscribe(ch <-chan uint64) {
	_m.Called(ch)
}

type mockConstructorTestingTNewBlockHeightWatcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlockHeightWatcher creates a new instance of BlockHeightWatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlockHeightWatcher(t mockConstructorTestingTNewBlockHeightWatcher) *BlockHeightWatcher {
	mock := &BlockHeightWatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}