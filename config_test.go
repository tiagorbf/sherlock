package main

import (
	"testing"

	"github.com/trivago/tgo/ttesting"
)

var application1 = ApplicationAttributes{Name: "app1", CloneUrl: "cloneUrl1", Path: "path1", Dependencies: []string{"app2", "app3", "app4"}}
var application2 = ApplicationAttributes{Name: "app2", CloneUrl: "cloneUrl2", Path: "path2", Dependencies: []string{"app3"}}
var application3 = ApplicationAttributes{Name: "app3", CloneUrl: "cloneUrl3", Path: "path3", Dependencies: []string{}}
var application4 = ApplicationAttributes{Name: "app4", CloneUrl: "cloneUrl4", Path: "path4", Dependencies: []string{"app1"}}

var appsConfig = Applications{ApplicationsMap: []ApplicationAttributes{application1, application2, application3, application4}}

func TestGetApp1Dependencies(t *testing.T) {
	expect := ttesting.NewExpect(t)
	expectedDependencies := appsConfig
	appDependencies := getAppDependencies(appsConfig, []string{"app1"}, []ApplicationAttributes{})

	expect.Equal(expectedDependencies.ApplicationsMap, appDependencies)
}

func TestGetApp2Dependencies(t *testing.T) {
	expect := ttesting.NewExpect(t)
	expectedDependencies := Applications{ApplicationsMap: []ApplicationAttributes{application2, application3}}
	appDependencies := getAppDependencies(appsConfig, []string{"app2"}, []ApplicationAttributes{})

	expect.Equal(expectedDependencies.ApplicationsMap, appDependencies)
}

func TestGetApp3Dependencies(t *testing.T) {
	expect := ttesting.NewExpect(t)
	expectedDependencies := Applications{ApplicationsMap: []ApplicationAttributes{application3}}
	appDependencies := getAppDependencies(appsConfig, []string{"app3"}, []ApplicationAttributes{})

	expect.Equal(expectedDependencies.ApplicationsMap, appDependencies)
}

func TestGetApp4Dependencies(t *testing.T) {
	expect := ttesting.NewExpect(t)
	expectedDependencies := Applications{ApplicationsMap: []ApplicationAttributes{application4, application1, application2, application3}}
	appDependencies := getAppDependencies(appsConfig, []string{"app4"}, []ApplicationAttributes{})

	expect.Equal(expectedDependencies.ApplicationsMap, appDependencies)
}

func TestGetSliceOfAppNames(t *testing.T) {
	expect := ttesting.NewExpect(t)
	expectedSlice := []string{"app1", "app2"}
	expect.Equal(expectedSlice, getSliceOfAppNames([]ApplicationAttributes{application1, application2}))

	// Test when empty arry is given
	expectedSlice = []string{}
	expect.Equal(expectedSlice, getSliceOfAppNames([]ApplicationAttributes{}))
}

func TestSliceContainsString(t *testing.T) {
	expect := ttesting.NewExpect(t)
	expect.True(sliceContainsString([]string{"app1", "app2"}, "app1"))

	expect.False(sliceContainsString([]string{}, "app1"))
	expect.False(sliceContainsString([]string{}, ""))
	expect.False(sliceContainsString([]string{"app1", "app2"}, "app4"))
}
