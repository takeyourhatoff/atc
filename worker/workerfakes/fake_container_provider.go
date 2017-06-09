// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"os"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeContainerProvider struct {
	FindCreatedContainerByHandleStub        func(logger lager.Logger, handle string, teamID int) (worker.Container, bool, error)
	findCreatedContainerByHandleMutex       sync.RWMutex
	findCreatedContainerByHandleArgsForCall []struct {
		logger lager.Logger
		handle string
		teamID int
	}
	findCreatedContainerByHandleReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	findCreatedContainerByHandleReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	FindOrCreateContainerStub        func(logger lager.Logger, cancel <-chan os.Signal, resourceUser db.ResourceUser, owner db.ContainerOwner, delegate worker.ImageFetchingDelegate, metadata db.ContainerMetadata, spec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes) (worker.Container, error)
	findOrCreateContainerMutex       sync.RWMutex
	findOrCreateContainerArgsForCall []struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		resourceUser  db.ResourceUser
		owner         db.ContainerOwner
		delegate      worker.ImageFetchingDelegate
		metadata      db.ContainerMetadata
		spec          worker.ContainerSpec
		resourceTypes atc.VersionedResourceTypes
	}
	findOrCreateContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	findOrCreateContainerReturnsOnCall map[int]struct {
		result1 worker.Container
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerProvider) FindCreatedContainerByHandle(logger lager.Logger, handle string, teamID int) (worker.Container, bool, error) {
	fake.findCreatedContainerByHandleMutex.Lock()
	ret, specificReturn := fake.findCreatedContainerByHandleReturnsOnCall[len(fake.findCreatedContainerByHandleArgsForCall)]
	fake.findCreatedContainerByHandleArgsForCall = append(fake.findCreatedContainerByHandleArgsForCall, struct {
		logger lager.Logger
		handle string
		teamID int
	}{logger, handle, teamID})
	fake.recordInvocation("FindCreatedContainerByHandle", []interface{}{logger, handle, teamID})
	fake.findCreatedContainerByHandleMutex.Unlock()
	if fake.FindCreatedContainerByHandleStub != nil {
		return fake.FindCreatedContainerByHandleStub(logger, handle, teamID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findCreatedContainerByHandleReturns.result1, fake.findCreatedContainerByHandleReturns.result2, fake.findCreatedContainerByHandleReturns.result3
}

func (fake *FakeContainerProvider) FindCreatedContainerByHandleCallCount() int {
	fake.findCreatedContainerByHandleMutex.RLock()
	defer fake.findCreatedContainerByHandleMutex.RUnlock()
	return len(fake.findCreatedContainerByHandleArgsForCall)
}

func (fake *FakeContainerProvider) FindCreatedContainerByHandleArgsForCall(i int) (lager.Logger, string, int) {
	fake.findCreatedContainerByHandleMutex.RLock()
	defer fake.findCreatedContainerByHandleMutex.RUnlock()
	return fake.findCreatedContainerByHandleArgsForCall[i].logger, fake.findCreatedContainerByHandleArgsForCall[i].handle, fake.findCreatedContainerByHandleArgsForCall[i].teamID
}

func (fake *FakeContainerProvider) FindCreatedContainerByHandleReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.FindCreatedContainerByHandleStub = nil
	fake.findCreatedContainerByHandleReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerProvider) FindCreatedContainerByHandleReturnsOnCall(i int, result1 worker.Container, result2 bool, result3 error) {
	fake.FindCreatedContainerByHandleStub = nil
	if fake.findCreatedContainerByHandleReturnsOnCall == nil {
		fake.findCreatedContainerByHandleReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 bool
			result3 error
		})
	}
	fake.findCreatedContainerByHandleReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerProvider) FindOrCreateContainer(logger lager.Logger, cancel <-chan os.Signal, resourceUser db.ResourceUser, owner db.ContainerOwner, delegate worker.ImageFetchingDelegate, metadata db.ContainerMetadata, spec worker.ContainerSpec, resourceTypes atc.VersionedResourceTypes) (worker.Container, error) {
	fake.findOrCreateContainerMutex.Lock()
	ret, specificReturn := fake.findOrCreateContainerReturnsOnCall[len(fake.findOrCreateContainerArgsForCall)]
	fake.findOrCreateContainerArgsForCall = append(fake.findOrCreateContainerArgsForCall, struct {
		logger        lager.Logger
		cancel        <-chan os.Signal
		resourceUser  db.ResourceUser
		owner         db.ContainerOwner
		delegate      worker.ImageFetchingDelegate
		metadata      db.ContainerMetadata
		spec          worker.ContainerSpec
		resourceTypes atc.VersionedResourceTypes
	}{logger, cancel, resourceUser, owner, delegate, metadata, spec, resourceTypes})
	fake.recordInvocation("FindOrCreateContainer", []interface{}{logger, cancel, resourceUser, owner, delegate, metadata, spec, resourceTypes})
	fake.findOrCreateContainerMutex.Unlock()
	if fake.FindOrCreateContainerStub != nil {
		return fake.FindOrCreateContainerStub(logger, cancel, resourceUser, owner, delegate, metadata, spec, resourceTypes)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.findOrCreateContainerReturns.result1, fake.findOrCreateContainerReturns.result2
}

func (fake *FakeContainerProvider) FindOrCreateContainerCallCount() int {
	fake.findOrCreateContainerMutex.RLock()
	defer fake.findOrCreateContainerMutex.RUnlock()
	return len(fake.findOrCreateContainerArgsForCall)
}

func (fake *FakeContainerProvider) FindOrCreateContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, db.ResourceUser, db.ContainerOwner, worker.ImageFetchingDelegate, db.ContainerMetadata, worker.ContainerSpec, atc.VersionedResourceTypes) {
	fake.findOrCreateContainerMutex.RLock()
	defer fake.findOrCreateContainerMutex.RUnlock()
	return fake.findOrCreateContainerArgsForCall[i].logger, fake.findOrCreateContainerArgsForCall[i].cancel, fake.findOrCreateContainerArgsForCall[i].resourceUser, fake.findOrCreateContainerArgsForCall[i].owner, fake.findOrCreateContainerArgsForCall[i].delegate, fake.findOrCreateContainerArgsForCall[i].metadata, fake.findOrCreateContainerArgsForCall[i].spec, fake.findOrCreateContainerArgsForCall[i].resourceTypes
}

func (fake *FakeContainerProvider) FindOrCreateContainerReturns(result1 worker.Container, result2 error) {
	fake.FindOrCreateContainerStub = nil
	fake.findOrCreateContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerProvider) FindOrCreateContainerReturnsOnCall(i int, result1 worker.Container, result2 error) {
	fake.FindOrCreateContainerStub = nil
	if fake.findOrCreateContainerReturnsOnCall == nil {
		fake.findOrCreateContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Container
			result2 error
		})
	}
	fake.findOrCreateContainerReturnsOnCall[i] = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findCreatedContainerByHandleMutex.RLock()
	defer fake.findCreatedContainerByHandleMutex.RUnlock()
	fake.findOrCreateContainerMutex.RLock()
	defer fake.findOrCreateContainerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerProvider) recordInvocation(key string, args []interface{}) {
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

var _ worker.ContainerProvider = new(FakeContainerProvider)
