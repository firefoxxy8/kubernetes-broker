// Automatically generated by MockGen. DO NOT EDIT!
// Source: consul/consul_service.go

package consul

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of ConsulService interface
type MockConsulService struct {
	ctrl     *gomock.Controller
	recorder *_MockConsulServiceRecorder
}

// Recorder for MockConsulService (not exported)
type _MockConsulServiceRecorder struct {
	mock *MockConsulService
}

func NewMockConsulService(ctrl *gomock.Controller) *MockConsulService {
	mock := &MockConsulService{ctrl: ctrl}
	mock.recorder = &_MockConsulServiceRecorder{mock}
	return mock
}

func (_m *MockConsulService) EXPECT() *_MockConsulServiceRecorder {
	return _m.recorder
}

func (_m *MockConsulService) UpdateServiceTag(params []ConsulServiceParams, endpoint string) error {
	ret := _m.ctrl.Call(_m, "UpdateServiceTag", params, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConsulServiceRecorder) UpdateServiceTag(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateServiceTag", arg0, arg1)
}

func (_m *MockConsulService) GetServicesListWithPublicTagStatus(endpoint string) (map[string]bool, error) {
	ret := _m.ctrl.Call(_m, "GetServicesListWithPublicTagStatus", endpoint)
	ret0, _ := ret[0].(map[string]bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConsulServiceRecorder) GetServicesListWithPublicTagStatus(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetServicesListWithPublicTagStatus", arg0)
}