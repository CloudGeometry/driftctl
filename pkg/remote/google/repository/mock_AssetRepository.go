// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	mock "github.com/stretchr/testify/mock"
	asset "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

// MockAssetRepository is an autogenerated mock type for the AssetRepository type
type MockAssetRepository struct {
	mock.Mock
}

// SearchAllAddresses provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllAddresses() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllBuckets provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllBuckets() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllDNSManagedZones provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllDNSManagedZones() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllDatasets provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllDatasets() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllFirewalls provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllFirewalls() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllFunctions provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllFunctions() ([]*asset.Asset, error) {
	ret := _m.Called()

	var r0 []*asset.Asset
	if rf, ok := ret.Get(0).(func() []*asset.Asset); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.Asset)
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

// SearchAllInstanceGroups provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllInstanceGroups() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllInstances provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllInstances() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllNetworks provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllNetworks() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllRouters provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllRouters() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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

// SearchAllTables provides a mock function with given fields:
func (_m *MockAssetRepository) SearchAllTables() ([]*asset.ResourceSearchResult, error) {
	ret := _m.Called()

	var r0 []*asset.ResourceSearchResult
	if rf, ok := ret.Get(0).(func() []*asset.ResourceSearchResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*asset.ResourceSearchResult)
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
