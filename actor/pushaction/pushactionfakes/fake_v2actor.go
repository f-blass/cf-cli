// Code generated by counterfeiter. DO NOT EDIT.
package pushactionfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/actor/v2action"
)

type FakeV2Actor struct {
	BindRouteToApplicationStub        func(routeGUID string, appGUID string) (v2action.Warnings, error)
	bindRouteToApplicationMutex       sync.RWMutex
	bindRouteToApplicationArgsForCall []struct {
		routeGUID string
		appGUID   string
	}
	bindRouteToApplicationReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	bindRouteToApplicationReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	CreateApplicationStub        func(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	createApplicationMutex       sync.RWMutex
	createApplicationArgsForCall []struct {
		application v2action.Application
	}
	createApplicationReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	createApplicationReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	CreateRouteStub        func(route v2action.Route, generatePort bool) (v2action.Route, v2action.Warnings, error)
	createRouteMutex       sync.RWMutex
	createRouteArgsForCall []struct {
		route        v2action.Route
		generatePort bool
	}
	createRouteReturns struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	createRouteReturnsOnCall map[int]struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	FindRouteBoundToSpaceWithSettingsStub        func(route v2action.Route) (v2action.Route, v2action.Warnings, error)
	findRouteBoundToSpaceWithSettingsMutex       sync.RWMutex
	findRouteBoundToSpaceWithSettingsArgsForCall []struct {
		route v2action.Route
	}
	findRouteBoundToSpaceWithSettingsReturns struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	findRouteBoundToSpaceWithSettingsReturnsOnCall map[int]struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	GatherDirectoryResourcesStub        func(sourceDir string) ([]v2action.Resource, error)
	gatherDirectoryResourcesMutex       sync.RWMutex
	gatherDirectoryResourcesArgsForCall []struct {
		sourceDir string
	}
	gatherDirectoryResourcesReturns struct {
		result1 []v2action.Resource
		result2 error
	}
	gatherDirectoryResourcesReturnsOnCall map[int]struct {
		result1 []v2action.Resource
		result2 error
	}
	GetApplicationByNameAndSpaceStub        func(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		name      string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationRoutesStub        func(applicationGUID string) ([]v2action.Route, v2action.Warnings, error)
	getApplicationRoutesMutex       sync.RWMutex
	getApplicationRoutesArgsForCall []struct {
		applicationGUID string
	}
	getApplicationRoutesReturns struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	getApplicationRoutesReturnsOnCall map[int]struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}
	GetOrganizationDomainsStub        func(orgGUID string) ([]v2action.Domain, v2action.Warnings, error)
	getOrganizationDomainsMutex       sync.RWMutex
	getOrganizationDomainsArgsForCall []struct {
		orgGUID string
	}
	getOrganizationDomainsReturns struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationDomainsReturnsOnCall map[int]struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}
	PollJobStub        func(job v2action.Job) (v2action.Warnings, error)
	pollJobMutex       sync.RWMutex
	pollJobArgsForCall []struct {
		job v2action.Job
	}
	pollJobReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	pollJobReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	UpdateApplicationStub        func(application v2action.Application) (v2action.Application, v2action.Warnings, error)
	updateApplicationMutex       sync.RWMutex
	updateApplicationArgsForCall []struct {
		application v2action.Application
	}
	updateApplicationReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	updateApplicationReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	UploadApplicationPackageStub        func(appGUID string, existingResources []v2action.Resource, newResources io.Reader, newResourcesLength int64) (v2action.Job, v2action.Warnings, error)
	uploadApplicationPackageMutex       sync.RWMutex
	uploadApplicationPackageArgsForCall []struct {
		appGUID            string
		existingResources  []v2action.Resource
		newResources       io.Reader
		newResourcesLength int64
	}
	uploadApplicationPackageReturns struct {
		result1 v2action.Job
		result2 v2action.Warnings
		result3 error
	}
	uploadApplicationPackageReturnsOnCall map[int]struct {
		result1 v2action.Job
		result2 v2action.Warnings
		result3 error
	}
	ZipDirectoryResourcesStub        func(sourceDir string, filesToInclude []v2action.Resource) (string, error)
	zipDirectoryResourcesMutex       sync.RWMutex
	zipDirectoryResourcesArgsForCall []struct {
		sourceDir      string
		filesToInclude []v2action.Resource
	}
	zipDirectoryResourcesReturns struct {
		result1 string
		result2 error
	}
	zipDirectoryResourcesReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2Actor) BindRouteToApplication(routeGUID string, appGUID string) (v2action.Warnings, error) {
	fake.bindRouteToApplicationMutex.Lock()
	ret, specificReturn := fake.bindRouteToApplicationReturnsOnCall[len(fake.bindRouteToApplicationArgsForCall)]
	fake.bindRouteToApplicationArgsForCall = append(fake.bindRouteToApplicationArgsForCall, struct {
		routeGUID string
		appGUID   string
	}{routeGUID, appGUID})
	fake.recordInvocation("BindRouteToApplication", []interface{}{routeGUID, appGUID})
	fake.bindRouteToApplicationMutex.Unlock()
	if fake.BindRouteToApplicationStub != nil {
		return fake.BindRouteToApplicationStub(routeGUID, appGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindRouteToApplicationReturns.result1, fake.bindRouteToApplicationReturns.result2
}

func (fake *FakeV2Actor) BindRouteToApplicationCallCount() int {
	fake.bindRouteToApplicationMutex.RLock()
	defer fake.bindRouteToApplicationMutex.RUnlock()
	return len(fake.bindRouteToApplicationArgsForCall)
}

func (fake *FakeV2Actor) BindRouteToApplicationArgsForCall(i int) (string, string) {
	fake.bindRouteToApplicationMutex.RLock()
	defer fake.bindRouteToApplicationMutex.RUnlock()
	return fake.bindRouteToApplicationArgsForCall[i].routeGUID, fake.bindRouteToApplicationArgsForCall[i].appGUID
}

func (fake *FakeV2Actor) BindRouteToApplicationReturns(result1 v2action.Warnings, result2 error) {
	fake.BindRouteToApplicationStub = nil
	fake.bindRouteToApplicationReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) BindRouteToApplicationReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.BindRouteToApplicationStub = nil
	if fake.bindRouteToApplicationReturnsOnCall == nil {
		fake.bindRouteToApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.bindRouteToApplicationReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) CreateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error) {
	fake.createApplicationMutex.Lock()
	ret, specificReturn := fake.createApplicationReturnsOnCall[len(fake.createApplicationArgsForCall)]
	fake.createApplicationArgsForCall = append(fake.createApplicationArgsForCall, struct {
		application v2action.Application
	}{application})
	fake.recordInvocation("CreateApplication", []interface{}{application})
	fake.createApplicationMutex.Unlock()
	if fake.CreateApplicationStub != nil {
		return fake.CreateApplicationStub(application)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createApplicationReturns.result1, fake.createApplicationReturns.result2, fake.createApplicationReturns.result3
}

func (fake *FakeV2Actor) CreateApplicationCallCount() int {
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	return len(fake.createApplicationArgsForCall)
}

func (fake *FakeV2Actor) CreateApplicationArgsForCall(i int) v2action.Application {
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	return fake.createApplicationArgsForCall[i].application
}

func (fake *FakeV2Actor) CreateApplicationReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationStub = nil
	fake.createApplicationReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateApplicationReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.CreateApplicationStub = nil
	if fake.createApplicationReturnsOnCall == nil {
		fake.createApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createApplicationReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateRoute(route v2action.Route, generatePort bool) (v2action.Route, v2action.Warnings, error) {
	fake.createRouteMutex.Lock()
	ret, specificReturn := fake.createRouteReturnsOnCall[len(fake.createRouteArgsForCall)]
	fake.createRouteArgsForCall = append(fake.createRouteArgsForCall, struct {
		route        v2action.Route
		generatePort bool
	}{route, generatePort})
	fake.recordInvocation("CreateRoute", []interface{}{route, generatePort})
	fake.createRouteMutex.Unlock()
	if fake.CreateRouteStub != nil {
		return fake.CreateRouteStub(route, generatePort)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createRouteReturns.result1, fake.createRouteReturns.result2, fake.createRouteReturns.result3
}

func (fake *FakeV2Actor) CreateRouteCallCount() int {
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	return len(fake.createRouteArgsForCall)
}

func (fake *FakeV2Actor) CreateRouteArgsForCall(i int) (v2action.Route, bool) {
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	return fake.createRouteArgsForCall[i].route, fake.createRouteArgsForCall[i].generatePort
}

func (fake *FakeV2Actor) CreateRouteReturns(result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.CreateRouteStub = nil
	fake.createRouteReturns = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) CreateRouteReturnsOnCall(i int, result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.CreateRouteStub = nil
	if fake.createRouteReturnsOnCall == nil {
		fake.createRouteReturnsOnCall = make(map[int]struct {
			result1 v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.createRouteReturnsOnCall[i] = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) FindRouteBoundToSpaceWithSettings(route v2action.Route) (v2action.Route, v2action.Warnings, error) {
	fake.findRouteBoundToSpaceWithSettingsMutex.Lock()
	ret, specificReturn := fake.findRouteBoundToSpaceWithSettingsReturnsOnCall[len(fake.findRouteBoundToSpaceWithSettingsArgsForCall)]
	fake.findRouteBoundToSpaceWithSettingsArgsForCall = append(fake.findRouteBoundToSpaceWithSettingsArgsForCall, struct {
		route v2action.Route
	}{route})
	fake.recordInvocation("FindRouteBoundToSpaceWithSettings", []interface{}{route})
	fake.findRouteBoundToSpaceWithSettingsMutex.Unlock()
	if fake.FindRouteBoundToSpaceWithSettingsStub != nil {
		return fake.FindRouteBoundToSpaceWithSettingsStub(route)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findRouteBoundToSpaceWithSettingsReturns.result1, fake.findRouteBoundToSpaceWithSettingsReturns.result2, fake.findRouteBoundToSpaceWithSettingsReturns.result3
}

func (fake *FakeV2Actor) FindRouteBoundToSpaceWithSettingsCallCount() int {
	fake.findRouteBoundToSpaceWithSettingsMutex.RLock()
	defer fake.findRouteBoundToSpaceWithSettingsMutex.RUnlock()
	return len(fake.findRouteBoundToSpaceWithSettingsArgsForCall)
}

func (fake *FakeV2Actor) FindRouteBoundToSpaceWithSettingsArgsForCall(i int) v2action.Route {
	fake.findRouteBoundToSpaceWithSettingsMutex.RLock()
	defer fake.findRouteBoundToSpaceWithSettingsMutex.RUnlock()
	return fake.findRouteBoundToSpaceWithSettingsArgsForCall[i].route
}

func (fake *FakeV2Actor) FindRouteBoundToSpaceWithSettingsReturns(result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.FindRouteBoundToSpaceWithSettingsStub = nil
	fake.findRouteBoundToSpaceWithSettingsReturns = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) FindRouteBoundToSpaceWithSettingsReturnsOnCall(i int, result1 v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.FindRouteBoundToSpaceWithSettingsStub = nil
	if fake.findRouteBoundToSpaceWithSettingsReturnsOnCall == nil {
		fake.findRouteBoundToSpaceWithSettingsReturnsOnCall = make(map[int]struct {
			result1 v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.findRouteBoundToSpaceWithSettingsReturnsOnCall[i] = struct {
		result1 v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GatherDirectoryResources(sourceDir string) ([]v2action.Resource, error) {
	fake.gatherDirectoryResourcesMutex.Lock()
	ret, specificReturn := fake.gatherDirectoryResourcesReturnsOnCall[len(fake.gatherDirectoryResourcesArgsForCall)]
	fake.gatherDirectoryResourcesArgsForCall = append(fake.gatherDirectoryResourcesArgsForCall, struct {
		sourceDir string
	}{sourceDir})
	fake.recordInvocation("GatherDirectoryResources", []interface{}{sourceDir})
	fake.gatherDirectoryResourcesMutex.Unlock()
	if fake.GatherDirectoryResourcesStub != nil {
		return fake.GatherDirectoryResourcesStub(sourceDir)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.gatherDirectoryResourcesReturns.result1, fake.gatherDirectoryResourcesReturns.result2
}

func (fake *FakeV2Actor) GatherDirectoryResourcesCallCount() int {
	fake.gatherDirectoryResourcesMutex.RLock()
	defer fake.gatherDirectoryResourcesMutex.RUnlock()
	return len(fake.gatherDirectoryResourcesArgsForCall)
}

func (fake *FakeV2Actor) GatherDirectoryResourcesArgsForCall(i int) string {
	fake.gatherDirectoryResourcesMutex.RLock()
	defer fake.gatherDirectoryResourcesMutex.RUnlock()
	return fake.gatherDirectoryResourcesArgsForCall[i].sourceDir
}

func (fake *FakeV2Actor) GatherDirectoryResourcesReturns(result1 []v2action.Resource, result2 error) {
	fake.GatherDirectoryResourcesStub = nil
	fake.gatherDirectoryResourcesReturns = struct {
		result1 []v2action.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) GatherDirectoryResourcesReturnsOnCall(i int, result1 []v2action.Resource, result2 error) {
	fake.GatherDirectoryResourcesStub = nil
	if fake.gatherDirectoryResourcesReturnsOnCall == nil {
		fake.gatherDirectoryResourcesReturnsOnCall = make(map[int]struct {
			result1 []v2action.Resource
			result2 error
		})
	}
	fake.gatherDirectoryResourcesReturnsOnCall[i] = struct {
		result1 []v2action.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpace(name string, spaceGUID string) (v2action.Application, v2action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		name      string
		spaceGUID string
	}{name, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{name, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(name, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].name, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationRoutes(applicationGUID string) ([]v2action.Route, v2action.Warnings, error) {
	fake.getApplicationRoutesMutex.Lock()
	ret, specificReturn := fake.getApplicationRoutesReturnsOnCall[len(fake.getApplicationRoutesArgsForCall)]
	fake.getApplicationRoutesArgsForCall = append(fake.getApplicationRoutesArgsForCall, struct {
		applicationGUID string
	}{applicationGUID})
	fake.recordInvocation("GetApplicationRoutes", []interface{}{applicationGUID})
	fake.getApplicationRoutesMutex.Unlock()
	if fake.GetApplicationRoutesStub != nil {
		return fake.GetApplicationRoutesStub(applicationGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationRoutesReturns.result1, fake.getApplicationRoutesReturns.result2, fake.getApplicationRoutesReturns.result3
}

func (fake *FakeV2Actor) GetApplicationRoutesCallCount() int {
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	return len(fake.getApplicationRoutesArgsForCall)
}

func (fake *FakeV2Actor) GetApplicationRoutesArgsForCall(i int) string {
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	return fake.getApplicationRoutesArgsForCall[i].applicationGUID
}

func (fake *FakeV2Actor) GetApplicationRoutesReturns(result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationRoutesStub = nil
	fake.getApplicationRoutesReturns = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetApplicationRoutesReturnsOnCall(i int, result1 []v2action.Route, result2 v2action.Warnings, result3 error) {
	fake.GetApplicationRoutesStub = nil
	if fake.getApplicationRoutesReturnsOnCall == nil {
		fake.getApplicationRoutesReturnsOnCall = make(map[int]struct {
			result1 []v2action.Route
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationRoutesReturnsOnCall[i] = struct {
		result1 []v2action.Route
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetOrganizationDomains(orgGUID string) ([]v2action.Domain, v2action.Warnings, error) {
	fake.getOrganizationDomainsMutex.Lock()
	ret, specificReturn := fake.getOrganizationDomainsReturnsOnCall[len(fake.getOrganizationDomainsArgsForCall)]
	fake.getOrganizationDomainsArgsForCall = append(fake.getOrganizationDomainsArgsForCall, struct {
		orgGUID string
	}{orgGUID})
	fake.recordInvocation("GetOrganizationDomains", []interface{}{orgGUID})
	fake.getOrganizationDomainsMutex.Unlock()
	if fake.GetOrganizationDomainsStub != nil {
		return fake.GetOrganizationDomainsStub(orgGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getOrganizationDomainsReturns.result1, fake.getOrganizationDomainsReturns.result2, fake.getOrganizationDomainsReturns.result3
}

func (fake *FakeV2Actor) GetOrganizationDomainsCallCount() int {
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	return len(fake.getOrganizationDomainsArgsForCall)
}

func (fake *FakeV2Actor) GetOrganizationDomainsArgsForCall(i int) string {
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	return fake.getOrganizationDomainsArgsForCall[i].orgGUID
}

func (fake *FakeV2Actor) GetOrganizationDomainsReturns(result1 []v2action.Domain, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationDomainsStub = nil
	fake.getOrganizationDomainsReturns = struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) GetOrganizationDomainsReturnsOnCall(i int, result1 []v2action.Domain, result2 v2action.Warnings, result3 error) {
	fake.GetOrganizationDomainsStub = nil
	if fake.getOrganizationDomainsReturnsOnCall == nil {
		fake.getOrganizationDomainsReturnsOnCall = make(map[int]struct {
			result1 []v2action.Domain
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationDomainsReturnsOnCall[i] = struct {
		result1 []v2action.Domain
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) PollJob(job v2action.Job) (v2action.Warnings, error) {
	fake.pollJobMutex.Lock()
	ret, specificReturn := fake.pollJobReturnsOnCall[len(fake.pollJobArgsForCall)]
	fake.pollJobArgsForCall = append(fake.pollJobArgsForCall, struct {
		job v2action.Job
	}{job})
	fake.recordInvocation("PollJob", []interface{}{job})
	fake.pollJobMutex.Unlock()
	if fake.PollJobStub != nil {
		return fake.PollJobStub(job)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.pollJobReturns.result1, fake.pollJobReturns.result2
}

func (fake *FakeV2Actor) PollJobCallCount() int {
	fake.pollJobMutex.RLock()
	defer fake.pollJobMutex.RUnlock()
	return len(fake.pollJobArgsForCall)
}

func (fake *FakeV2Actor) PollJobArgsForCall(i int) v2action.Job {
	fake.pollJobMutex.RLock()
	defer fake.pollJobMutex.RUnlock()
	return fake.pollJobArgsForCall[i].job
}

func (fake *FakeV2Actor) PollJobReturns(result1 v2action.Warnings, result2 error) {
	fake.PollJobStub = nil
	fake.pollJobReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) PollJobReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.PollJobStub = nil
	if fake.pollJobReturnsOnCall == nil {
		fake.pollJobReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.pollJobReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) UpdateApplication(application v2action.Application) (v2action.Application, v2action.Warnings, error) {
	fake.updateApplicationMutex.Lock()
	ret, specificReturn := fake.updateApplicationReturnsOnCall[len(fake.updateApplicationArgsForCall)]
	fake.updateApplicationArgsForCall = append(fake.updateApplicationArgsForCall, struct {
		application v2action.Application
	}{application})
	fake.recordInvocation("UpdateApplication", []interface{}{application})
	fake.updateApplicationMutex.Unlock()
	if fake.UpdateApplicationStub != nil {
		return fake.UpdateApplicationStub(application)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.updateApplicationReturns.result1, fake.updateApplicationReturns.result2, fake.updateApplicationReturns.result3
}

func (fake *FakeV2Actor) UpdateApplicationCallCount() int {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return len(fake.updateApplicationArgsForCall)
}

func (fake *FakeV2Actor) UpdateApplicationArgsForCall(i int) v2action.Application {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return fake.updateApplicationArgsForCall[i].application
}

func (fake *FakeV2Actor) UpdateApplicationReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	fake.updateApplicationReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UpdateApplicationReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.UpdateApplicationStub = nil
	if fake.updateApplicationReturnsOnCall == nil {
		fake.updateApplicationReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.updateApplicationReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UploadApplicationPackage(appGUID string, existingResources []v2action.Resource, newResources io.Reader, newResourcesLength int64) (v2action.Job, v2action.Warnings, error) {
	var existingResourcesCopy []v2action.Resource
	if existingResources != nil {
		existingResourcesCopy = make([]v2action.Resource, len(existingResources))
		copy(existingResourcesCopy, existingResources)
	}
	fake.uploadApplicationPackageMutex.Lock()
	ret, specificReturn := fake.uploadApplicationPackageReturnsOnCall[len(fake.uploadApplicationPackageArgsForCall)]
	fake.uploadApplicationPackageArgsForCall = append(fake.uploadApplicationPackageArgsForCall, struct {
		appGUID            string
		existingResources  []v2action.Resource
		newResources       io.Reader
		newResourcesLength int64
	}{appGUID, existingResourcesCopy, newResources, newResourcesLength})
	fake.recordInvocation("UploadApplicationPackage", []interface{}{appGUID, existingResourcesCopy, newResources, newResourcesLength})
	fake.uploadApplicationPackageMutex.Unlock()
	if fake.UploadApplicationPackageStub != nil {
		return fake.UploadApplicationPackageStub(appGUID, existingResources, newResources, newResourcesLength)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.uploadApplicationPackageReturns.result1, fake.uploadApplicationPackageReturns.result2, fake.uploadApplicationPackageReturns.result3
}

func (fake *FakeV2Actor) UploadApplicationPackageCallCount() int {
	fake.uploadApplicationPackageMutex.RLock()
	defer fake.uploadApplicationPackageMutex.RUnlock()
	return len(fake.uploadApplicationPackageArgsForCall)
}

func (fake *FakeV2Actor) UploadApplicationPackageArgsForCall(i int) (string, []v2action.Resource, io.Reader, int64) {
	fake.uploadApplicationPackageMutex.RLock()
	defer fake.uploadApplicationPackageMutex.RUnlock()
	return fake.uploadApplicationPackageArgsForCall[i].appGUID, fake.uploadApplicationPackageArgsForCall[i].existingResources, fake.uploadApplicationPackageArgsForCall[i].newResources, fake.uploadApplicationPackageArgsForCall[i].newResourcesLength
}

func (fake *FakeV2Actor) UploadApplicationPackageReturns(result1 v2action.Job, result2 v2action.Warnings, result3 error) {
	fake.UploadApplicationPackageStub = nil
	fake.uploadApplicationPackageReturns = struct {
		result1 v2action.Job
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) UploadApplicationPackageReturnsOnCall(i int, result1 v2action.Job, result2 v2action.Warnings, result3 error) {
	fake.UploadApplicationPackageStub = nil
	if fake.uploadApplicationPackageReturnsOnCall == nil {
		fake.uploadApplicationPackageReturnsOnCall = make(map[int]struct {
			result1 v2action.Job
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.uploadApplicationPackageReturnsOnCall[i] = struct {
		result1 v2action.Job
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2Actor) ZipDirectoryResources(sourceDir string, filesToInclude []v2action.Resource) (string, error) {
	var filesToIncludeCopy []v2action.Resource
	if filesToInclude != nil {
		filesToIncludeCopy = make([]v2action.Resource, len(filesToInclude))
		copy(filesToIncludeCopy, filesToInclude)
	}
	fake.zipDirectoryResourcesMutex.Lock()
	ret, specificReturn := fake.zipDirectoryResourcesReturnsOnCall[len(fake.zipDirectoryResourcesArgsForCall)]
	fake.zipDirectoryResourcesArgsForCall = append(fake.zipDirectoryResourcesArgsForCall, struct {
		sourceDir      string
		filesToInclude []v2action.Resource
	}{sourceDir, filesToIncludeCopy})
	fake.recordInvocation("ZipDirectoryResources", []interface{}{sourceDir, filesToIncludeCopy})
	fake.zipDirectoryResourcesMutex.Unlock()
	if fake.ZipDirectoryResourcesStub != nil {
		return fake.ZipDirectoryResourcesStub(sourceDir, filesToInclude)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.zipDirectoryResourcesReturns.result1, fake.zipDirectoryResourcesReturns.result2
}

func (fake *FakeV2Actor) ZipDirectoryResourcesCallCount() int {
	fake.zipDirectoryResourcesMutex.RLock()
	defer fake.zipDirectoryResourcesMutex.RUnlock()
	return len(fake.zipDirectoryResourcesArgsForCall)
}

func (fake *FakeV2Actor) ZipDirectoryResourcesArgsForCall(i int) (string, []v2action.Resource) {
	fake.zipDirectoryResourcesMutex.RLock()
	defer fake.zipDirectoryResourcesMutex.RUnlock()
	return fake.zipDirectoryResourcesArgsForCall[i].sourceDir, fake.zipDirectoryResourcesArgsForCall[i].filesToInclude
}

func (fake *FakeV2Actor) ZipDirectoryResourcesReturns(result1 string, result2 error) {
	fake.ZipDirectoryResourcesStub = nil
	fake.zipDirectoryResourcesReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) ZipDirectoryResourcesReturnsOnCall(i int, result1 string, result2 error) {
	fake.ZipDirectoryResourcesStub = nil
	if fake.zipDirectoryResourcesReturnsOnCall == nil {
		fake.zipDirectoryResourcesReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.zipDirectoryResourcesReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeV2Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindRouteToApplicationMutex.RLock()
	defer fake.bindRouteToApplicationMutex.RUnlock()
	fake.createApplicationMutex.RLock()
	defer fake.createApplicationMutex.RUnlock()
	fake.createRouteMutex.RLock()
	defer fake.createRouteMutex.RUnlock()
	fake.findRouteBoundToSpaceWithSettingsMutex.RLock()
	defer fake.findRouteBoundToSpaceWithSettingsMutex.RUnlock()
	fake.gatherDirectoryResourcesMutex.RLock()
	defer fake.gatherDirectoryResourcesMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationRoutesMutex.RLock()
	defer fake.getApplicationRoutesMutex.RUnlock()
	fake.getOrganizationDomainsMutex.RLock()
	defer fake.getOrganizationDomainsMutex.RUnlock()
	fake.pollJobMutex.RLock()
	defer fake.pollJobMutex.RUnlock()
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	fake.uploadApplicationPackageMutex.RLock()
	defer fake.uploadApplicationPackageMutex.RUnlock()
	fake.zipDirectoryResourcesMutex.RLock()
	defer fake.zipDirectoryResourcesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV2Actor) recordInvocation(key string, args []interface{}) {
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

var _ pushaction.V2Actor = new(FakeV2Actor)
