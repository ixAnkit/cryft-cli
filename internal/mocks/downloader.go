// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Downloader is an autogenerated mock type for the Downloader type
type Downloader struct {
	mock.Mock
}

// Download provides a mock function with given fields: url
func (_m *Downloader) Download(url string) ([]byte, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReleasesForRepo provides a mock function with given fields: org, repo
func (_m *Downloader) GetAllReleasesForRepo(org string, repo string) ([]string, error) {
	ret := _m.Called(org, repo)

	if len(ret) == 0 {
		panic("no return value specified for GetAllReleasesForRepo")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, error)); ok {
		return rf(org, repo)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(org, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(org, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestPreReleaseVersion provides a mock function with given fields: org, repo
func (_m *Downloader) GetLatestPreReleaseVersion(org string, repo string) (string, error) {
	ret := _m.Called(org, repo)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestPreReleaseVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(org, repo)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(org, repo)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(org, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestReleaseVersion provides a mock function with given fields: releaseURL
func (_m *Downloader) GetLatestReleaseVersion(releaseURL string) (string, error) {
	ret := _m.Called(releaseURL)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestReleaseVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(releaseURL)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(releaseURL)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(releaseURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDownloader creates a new instance of Downloader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDownloader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Downloader {
	mock := &Downloader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
