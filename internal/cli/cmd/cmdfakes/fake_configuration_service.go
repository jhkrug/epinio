// Copyright © 2021 - 2023 SUSE LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/epinio/epinio/internal/cli/cmd"
	"github.com/epinio/epinio/internal/cli/usercmd"
)

type FakeConfigurationService struct {
	AppsMatchingStub        func(string) []string
	appsMatchingMutex       sync.RWMutex
	appsMatchingArgsForCall []struct {
		arg1 string
	}
	appsMatchingReturns struct {
		result1 []string
	}
	appsMatchingReturnsOnCall map[int]struct {
		result1 []string
	}
	BindConfigurationStub        func(string, string) error
	bindConfigurationMutex       sync.RWMutex
	bindConfigurationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	bindConfigurationReturns struct {
		result1 error
	}
	bindConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigurationDetailsStub        func(string) error
	configurationDetailsMutex       sync.RWMutex
	configurationDetailsArgsForCall []struct {
		arg1 string
	}
	configurationDetailsReturns struct {
		result1 error
	}
	configurationDetailsReturnsOnCall map[int]struct {
		result1 error
	}
	ConfigurationMatchingStub        func(string) []string
	configurationMatchingMutex       sync.RWMutex
	configurationMatchingArgsForCall []struct {
		arg1 string
	}
	configurationMatchingReturns struct {
		result1 []string
	}
	configurationMatchingReturnsOnCall map[int]struct {
		result1 []string
	}
	ConfigurationsStub        func(bool) error
	configurationsMutex       sync.RWMutex
	configurationsArgsForCall []struct {
		arg1 bool
	}
	configurationsReturns struct {
		result1 error
	}
	configurationsReturnsOnCall map[int]struct {
		result1 error
	}
	CreateConfigurationStub        func(string, []string) error
	createConfigurationMutex       sync.RWMutex
	createConfigurationArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	createConfigurationReturns struct {
		result1 error
	}
	createConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteConfigurationStub        func([]string, bool, bool) error
	deleteConfigurationMutex       sync.RWMutex
	deleteConfigurationArgsForCall []struct {
		arg1 []string
		arg2 bool
		arg3 bool
	}
	deleteConfigurationReturns struct {
		result1 error
	}
	deleteConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	GetAPIStub        func() usercmd.APIClient
	getAPIMutex       sync.RWMutex
	getAPIArgsForCall []struct {
	}
	getAPIReturns struct {
		result1 usercmd.APIClient
	}
	getAPIReturnsOnCall map[int]struct {
		result1 usercmd.APIClient
	}
	UnbindConfigurationStub        func(string, string) error
	unbindConfigurationMutex       sync.RWMutex
	unbindConfigurationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	unbindConfigurationReturns struct {
		result1 error
	}
	unbindConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateConfigurationStub        func(string, []string, map[string]string) error
	updateConfigurationMutex       sync.RWMutex
	updateConfigurationArgsForCall []struct {
		arg1 string
		arg2 []string
		arg3 map[string]string
	}
	updateConfigurationReturns struct {
		result1 error
	}
	updateConfigurationReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigurationService) AppsMatching(arg1 string) []string {
	fake.appsMatchingMutex.Lock()
	ret, specificReturn := fake.appsMatchingReturnsOnCall[len(fake.appsMatchingArgsForCall)]
	fake.appsMatchingArgsForCall = append(fake.appsMatchingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.AppsMatchingStub
	fakeReturns := fake.appsMatchingReturns
	fake.recordInvocation("AppsMatching", []interface{}{arg1})
	fake.appsMatchingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) AppsMatchingCallCount() int {
	fake.appsMatchingMutex.RLock()
	defer fake.appsMatchingMutex.RUnlock()
	return len(fake.appsMatchingArgsForCall)
}

func (fake *FakeConfigurationService) AppsMatchingCalls(stub func(string) []string) {
	fake.appsMatchingMutex.Lock()
	defer fake.appsMatchingMutex.Unlock()
	fake.AppsMatchingStub = stub
}

func (fake *FakeConfigurationService) AppsMatchingArgsForCall(i int) string {
	fake.appsMatchingMutex.RLock()
	defer fake.appsMatchingMutex.RUnlock()
	argsForCall := fake.appsMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigurationService) AppsMatchingReturns(result1 []string) {
	fake.appsMatchingMutex.Lock()
	defer fake.appsMatchingMutex.Unlock()
	fake.AppsMatchingStub = nil
	fake.appsMatchingReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeConfigurationService) AppsMatchingReturnsOnCall(i int, result1 []string) {
	fake.appsMatchingMutex.Lock()
	defer fake.appsMatchingMutex.Unlock()
	fake.AppsMatchingStub = nil
	if fake.appsMatchingReturnsOnCall == nil {
		fake.appsMatchingReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.appsMatchingReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeConfigurationService) BindConfiguration(arg1 string, arg2 string) error {
	fake.bindConfigurationMutex.Lock()
	ret, specificReturn := fake.bindConfigurationReturnsOnCall[len(fake.bindConfigurationArgsForCall)]
	fake.bindConfigurationArgsForCall = append(fake.bindConfigurationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.BindConfigurationStub
	fakeReturns := fake.bindConfigurationReturns
	fake.recordInvocation("BindConfiguration", []interface{}{arg1, arg2})
	fake.bindConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) BindConfigurationCallCount() int {
	fake.bindConfigurationMutex.RLock()
	defer fake.bindConfigurationMutex.RUnlock()
	return len(fake.bindConfigurationArgsForCall)
}

func (fake *FakeConfigurationService) BindConfigurationCalls(stub func(string, string) error) {
	fake.bindConfigurationMutex.Lock()
	defer fake.bindConfigurationMutex.Unlock()
	fake.BindConfigurationStub = stub
}

func (fake *FakeConfigurationService) BindConfigurationArgsForCall(i int) (string, string) {
	fake.bindConfigurationMutex.RLock()
	defer fake.bindConfigurationMutex.RUnlock()
	argsForCall := fake.bindConfigurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigurationService) BindConfigurationReturns(result1 error) {
	fake.bindConfigurationMutex.Lock()
	defer fake.bindConfigurationMutex.Unlock()
	fake.BindConfigurationStub = nil
	fake.bindConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) BindConfigurationReturnsOnCall(i int, result1 error) {
	fake.bindConfigurationMutex.Lock()
	defer fake.bindConfigurationMutex.Unlock()
	fake.BindConfigurationStub = nil
	if fake.bindConfigurationReturnsOnCall == nil {
		fake.bindConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bindConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) ConfigurationDetails(arg1 string) error {
	fake.configurationDetailsMutex.Lock()
	ret, specificReturn := fake.configurationDetailsReturnsOnCall[len(fake.configurationDetailsArgsForCall)]
	fake.configurationDetailsArgsForCall = append(fake.configurationDetailsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ConfigurationDetailsStub
	fakeReturns := fake.configurationDetailsReturns
	fake.recordInvocation("ConfigurationDetails", []interface{}{arg1})
	fake.configurationDetailsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) ConfigurationDetailsCallCount() int {
	fake.configurationDetailsMutex.RLock()
	defer fake.configurationDetailsMutex.RUnlock()
	return len(fake.configurationDetailsArgsForCall)
}

func (fake *FakeConfigurationService) ConfigurationDetailsCalls(stub func(string) error) {
	fake.configurationDetailsMutex.Lock()
	defer fake.configurationDetailsMutex.Unlock()
	fake.ConfigurationDetailsStub = stub
}

func (fake *FakeConfigurationService) ConfigurationDetailsArgsForCall(i int) string {
	fake.configurationDetailsMutex.RLock()
	defer fake.configurationDetailsMutex.RUnlock()
	argsForCall := fake.configurationDetailsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigurationService) ConfigurationDetailsReturns(result1 error) {
	fake.configurationDetailsMutex.Lock()
	defer fake.configurationDetailsMutex.Unlock()
	fake.ConfigurationDetailsStub = nil
	fake.configurationDetailsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) ConfigurationDetailsReturnsOnCall(i int, result1 error) {
	fake.configurationDetailsMutex.Lock()
	defer fake.configurationDetailsMutex.Unlock()
	fake.ConfigurationDetailsStub = nil
	if fake.configurationDetailsReturnsOnCall == nil {
		fake.configurationDetailsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configurationDetailsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) ConfigurationMatching(arg1 string) []string {
	fake.configurationMatchingMutex.Lock()
	ret, specificReturn := fake.configurationMatchingReturnsOnCall[len(fake.configurationMatchingArgsForCall)]
	fake.configurationMatchingArgsForCall = append(fake.configurationMatchingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ConfigurationMatchingStub
	fakeReturns := fake.configurationMatchingReturns
	fake.recordInvocation("ConfigurationMatching", []interface{}{arg1})
	fake.configurationMatchingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) ConfigurationMatchingCallCount() int {
	fake.configurationMatchingMutex.RLock()
	defer fake.configurationMatchingMutex.RUnlock()
	return len(fake.configurationMatchingArgsForCall)
}

func (fake *FakeConfigurationService) ConfigurationMatchingCalls(stub func(string) []string) {
	fake.configurationMatchingMutex.Lock()
	defer fake.configurationMatchingMutex.Unlock()
	fake.ConfigurationMatchingStub = stub
}

func (fake *FakeConfigurationService) ConfigurationMatchingArgsForCall(i int) string {
	fake.configurationMatchingMutex.RLock()
	defer fake.configurationMatchingMutex.RUnlock()
	argsForCall := fake.configurationMatchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigurationService) ConfigurationMatchingReturns(result1 []string) {
	fake.configurationMatchingMutex.Lock()
	defer fake.configurationMatchingMutex.Unlock()
	fake.ConfigurationMatchingStub = nil
	fake.configurationMatchingReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeConfigurationService) ConfigurationMatchingReturnsOnCall(i int, result1 []string) {
	fake.configurationMatchingMutex.Lock()
	defer fake.configurationMatchingMutex.Unlock()
	fake.ConfigurationMatchingStub = nil
	if fake.configurationMatchingReturnsOnCall == nil {
		fake.configurationMatchingReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.configurationMatchingReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeConfigurationService) Configurations(arg1 bool) error {
	fake.configurationsMutex.Lock()
	ret, specificReturn := fake.configurationsReturnsOnCall[len(fake.configurationsArgsForCall)]
	fake.configurationsArgsForCall = append(fake.configurationsArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.ConfigurationsStub
	fakeReturns := fake.configurationsReturns
	fake.recordInvocation("Configurations", []interface{}{arg1})
	fake.configurationsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) ConfigurationsCallCount() int {
	fake.configurationsMutex.RLock()
	defer fake.configurationsMutex.RUnlock()
	return len(fake.configurationsArgsForCall)
}

func (fake *FakeConfigurationService) ConfigurationsCalls(stub func(bool) error) {
	fake.configurationsMutex.Lock()
	defer fake.configurationsMutex.Unlock()
	fake.ConfigurationsStub = stub
}

func (fake *FakeConfigurationService) ConfigurationsArgsForCall(i int) bool {
	fake.configurationsMutex.RLock()
	defer fake.configurationsMutex.RUnlock()
	argsForCall := fake.configurationsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConfigurationService) ConfigurationsReturns(result1 error) {
	fake.configurationsMutex.Lock()
	defer fake.configurationsMutex.Unlock()
	fake.ConfigurationsStub = nil
	fake.configurationsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) ConfigurationsReturnsOnCall(i int, result1 error) {
	fake.configurationsMutex.Lock()
	defer fake.configurationsMutex.Unlock()
	fake.ConfigurationsStub = nil
	if fake.configurationsReturnsOnCall == nil {
		fake.configurationsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configurationsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) CreateConfiguration(arg1 string, arg2 []string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.createConfigurationMutex.Lock()
	ret, specificReturn := fake.createConfigurationReturnsOnCall[len(fake.createConfigurationArgsForCall)]
	fake.createConfigurationArgsForCall = append(fake.createConfigurationArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.CreateConfigurationStub
	fakeReturns := fake.createConfigurationReturns
	fake.recordInvocation("CreateConfiguration", []interface{}{arg1, arg2Copy})
	fake.createConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) CreateConfigurationCallCount() int {
	fake.createConfigurationMutex.RLock()
	defer fake.createConfigurationMutex.RUnlock()
	return len(fake.createConfigurationArgsForCall)
}

func (fake *FakeConfigurationService) CreateConfigurationCalls(stub func(string, []string) error) {
	fake.createConfigurationMutex.Lock()
	defer fake.createConfigurationMutex.Unlock()
	fake.CreateConfigurationStub = stub
}

func (fake *FakeConfigurationService) CreateConfigurationArgsForCall(i int) (string, []string) {
	fake.createConfigurationMutex.RLock()
	defer fake.createConfigurationMutex.RUnlock()
	argsForCall := fake.createConfigurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigurationService) CreateConfigurationReturns(result1 error) {
	fake.createConfigurationMutex.Lock()
	defer fake.createConfigurationMutex.Unlock()
	fake.CreateConfigurationStub = nil
	fake.createConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) CreateConfigurationReturnsOnCall(i int, result1 error) {
	fake.createConfigurationMutex.Lock()
	defer fake.createConfigurationMutex.Unlock()
	fake.CreateConfigurationStub = nil
	if fake.createConfigurationReturnsOnCall == nil {
		fake.createConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) DeleteConfiguration(arg1 []string, arg2 bool, arg3 bool) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteConfigurationMutex.Lock()
	ret, specificReturn := fake.deleteConfigurationReturnsOnCall[len(fake.deleteConfigurationArgsForCall)]
	fake.deleteConfigurationArgsForCall = append(fake.deleteConfigurationArgsForCall, struct {
		arg1 []string
		arg2 bool
		arg3 bool
	}{arg1Copy, arg2, arg3})
	stub := fake.DeleteConfigurationStub
	fakeReturns := fake.deleteConfigurationReturns
	fake.recordInvocation("DeleteConfiguration", []interface{}{arg1Copy, arg2, arg3})
	fake.deleteConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) DeleteConfigurationCallCount() int {
	fake.deleteConfigurationMutex.RLock()
	defer fake.deleteConfigurationMutex.RUnlock()
	return len(fake.deleteConfigurationArgsForCall)
}

func (fake *FakeConfigurationService) DeleteConfigurationCalls(stub func([]string, bool, bool) error) {
	fake.deleteConfigurationMutex.Lock()
	defer fake.deleteConfigurationMutex.Unlock()
	fake.DeleteConfigurationStub = stub
}

func (fake *FakeConfigurationService) DeleteConfigurationArgsForCall(i int) ([]string, bool, bool) {
	fake.deleteConfigurationMutex.RLock()
	defer fake.deleteConfigurationMutex.RUnlock()
	argsForCall := fake.deleteConfigurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConfigurationService) DeleteConfigurationReturns(result1 error) {
	fake.deleteConfigurationMutex.Lock()
	defer fake.deleteConfigurationMutex.Unlock()
	fake.DeleteConfigurationStub = nil
	fake.deleteConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) DeleteConfigurationReturnsOnCall(i int, result1 error) {
	fake.deleteConfigurationMutex.Lock()
	defer fake.deleteConfigurationMutex.Unlock()
	fake.DeleteConfigurationStub = nil
	if fake.deleteConfigurationReturnsOnCall == nil {
		fake.deleteConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) GetAPI() usercmd.APIClient {
	fake.getAPIMutex.Lock()
	ret, specificReturn := fake.getAPIReturnsOnCall[len(fake.getAPIArgsForCall)]
	fake.getAPIArgsForCall = append(fake.getAPIArgsForCall, struct {
	}{})
	stub := fake.GetAPIStub
	fakeReturns := fake.getAPIReturns
	fake.recordInvocation("GetAPI", []interface{}{})
	fake.getAPIMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) GetAPICallCount() int {
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	return len(fake.getAPIArgsForCall)
}

func (fake *FakeConfigurationService) GetAPICalls(stub func() usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = stub
}

func (fake *FakeConfigurationService) GetAPIReturns(result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	fake.getAPIReturns = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeConfigurationService) GetAPIReturnsOnCall(i int, result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	if fake.getAPIReturnsOnCall == nil {
		fake.getAPIReturnsOnCall = make(map[int]struct {
			result1 usercmd.APIClient
		})
	}
	fake.getAPIReturnsOnCall[i] = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeConfigurationService) UnbindConfiguration(arg1 string, arg2 string) error {
	fake.unbindConfigurationMutex.Lock()
	ret, specificReturn := fake.unbindConfigurationReturnsOnCall[len(fake.unbindConfigurationArgsForCall)]
	fake.unbindConfigurationArgsForCall = append(fake.unbindConfigurationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.UnbindConfigurationStub
	fakeReturns := fake.unbindConfigurationReturns
	fake.recordInvocation("UnbindConfiguration", []interface{}{arg1, arg2})
	fake.unbindConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) UnbindConfigurationCallCount() int {
	fake.unbindConfigurationMutex.RLock()
	defer fake.unbindConfigurationMutex.RUnlock()
	return len(fake.unbindConfigurationArgsForCall)
}

func (fake *FakeConfigurationService) UnbindConfigurationCalls(stub func(string, string) error) {
	fake.unbindConfigurationMutex.Lock()
	defer fake.unbindConfigurationMutex.Unlock()
	fake.UnbindConfigurationStub = stub
}

func (fake *FakeConfigurationService) UnbindConfigurationArgsForCall(i int) (string, string) {
	fake.unbindConfigurationMutex.RLock()
	defer fake.unbindConfigurationMutex.RUnlock()
	argsForCall := fake.unbindConfigurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConfigurationService) UnbindConfigurationReturns(result1 error) {
	fake.unbindConfigurationMutex.Lock()
	defer fake.unbindConfigurationMutex.Unlock()
	fake.UnbindConfigurationStub = nil
	fake.unbindConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) UnbindConfigurationReturnsOnCall(i int, result1 error) {
	fake.unbindConfigurationMutex.Lock()
	defer fake.unbindConfigurationMutex.Unlock()
	fake.UnbindConfigurationStub = nil
	if fake.unbindConfigurationReturnsOnCall == nil {
		fake.unbindConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) UpdateConfiguration(arg1 string, arg2 []string, arg3 map[string]string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.updateConfigurationMutex.Lock()
	ret, specificReturn := fake.updateConfigurationReturnsOnCall[len(fake.updateConfigurationArgsForCall)]
	fake.updateConfigurationArgsForCall = append(fake.updateConfigurationArgsForCall, struct {
		arg1 string
		arg2 []string
		arg3 map[string]string
	}{arg1, arg2Copy, arg3})
	stub := fake.UpdateConfigurationStub
	fakeReturns := fake.updateConfigurationReturns
	fake.recordInvocation("UpdateConfiguration", []interface{}{arg1, arg2Copy, arg3})
	fake.updateConfigurationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConfigurationService) UpdateConfigurationCallCount() int {
	fake.updateConfigurationMutex.RLock()
	defer fake.updateConfigurationMutex.RUnlock()
	return len(fake.updateConfigurationArgsForCall)
}

func (fake *FakeConfigurationService) UpdateConfigurationCalls(stub func(string, []string, map[string]string) error) {
	fake.updateConfigurationMutex.Lock()
	defer fake.updateConfigurationMutex.Unlock()
	fake.UpdateConfigurationStub = stub
}

func (fake *FakeConfigurationService) UpdateConfigurationArgsForCall(i int) (string, []string, map[string]string) {
	fake.updateConfigurationMutex.RLock()
	defer fake.updateConfigurationMutex.RUnlock()
	argsForCall := fake.updateConfigurationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConfigurationService) UpdateConfigurationReturns(result1 error) {
	fake.updateConfigurationMutex.Lock()
	defer fake.updateConfigurationMutex.Unlock()
	fake.UpdateConfigurationStub = nil
	fake.updateConfigurationReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) UpdateConfigurationReturnsOnCall(i int, result1 error) {
	fake.updateConfigurationMutex.Lock()
	defer fake.updateConfigurationMutex.Unlock()
	fake.UpdateConfigurationStub = nil
	if fake.updateConfigurationReturnsOnCall == nil {
		fake.updateConfigurationReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateConfigurationReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConfigurationService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appsMatchingMutex.RLock()
	defer fake.appsMatchingMutex.RUnlock()
	fake.bindConfigurationMutex.RLock()
	defer fake.bindConfigurationMutex.RUnlock()
	fake.configurationDetailsMutex.RLock()
	defer fake.configurationDetailsMutex.RUnlock()
	fake.configurationMatchingMutex.RLock()
	defer fake.configurationMatchingMutex.RUnlock()
	fake.configurationsMutex.RLock()
	defer fake.configurationsMutex.RUnlock()
	fake.createConfigurationMutex.RLock()
	defer fake.createConfigurationMutex.RUnlock()
	fake.deleteConfigurationMutex.RLock()
	defer fake.deleteConfigurationMutex.RUnlock()
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	fake.unbindConfigurationMutex.RLock()
	defer fake.unbindConfigurationMutex.RUnlock()
	fake.updateConfigurationMutex.RLock()
	defer fake.updateConfigurationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfigurationService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ cmd.ConfigurationService = new(FakeConfigurationService)