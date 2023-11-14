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

type FakeServiceChartValueMatcher struct {
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceChartValueMatcher) GetAPI() usercmd.APIClient {
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

func (fake *FakeServiceChartValueMatcher) GetAPICallCount() int {
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	return len(fake.getAPIArgsForCall)
}

func (fake *FakeServiceChartValueMatcher) GetAPICalls(stub func() usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = stub
}

func (fake *FakeServiceChartValueMatcher) GetAPIReturns(result1 usercmd.APIClient) {
	fake.getAPIMutex.Lock()
	defer fake.getAPIMutex.Unlock()
	fake.GetAPIStub = nil
	fake.getAPIReturns = struct {
		result1 usercmd.APIClient
	}{result1}
}

func (fake *FakeServiceChartValueMatcher) GetAPIReturnsOnCall(i int, result1 usercmd.APIClient) {
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

func (fake *FakeServiceChartValueMatcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAPIMutex.RLock()
	defer fake.getAPIMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceChartValueMatcher) recordInvocation(key string, args []interface{}) {
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

var _ cmd.ServiceChartValueMatcher = new(FakeServiceChartValueMatcher)