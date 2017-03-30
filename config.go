package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

type Applications struct {
	ApplicationsMap []ApplicationAttributes `json:"applications"`
}

type ApplicationAttributes struct {
	Name         string   `json:"name"`
	CloneUrl     string   `json:"clone_url"`
	Path         string   `json:"path"`
	Dependencies []string `json:"dependencies"`
}

func initConfig() Applications {
	app := Applications{}
	fileName := "applications.json"
	pathToFile, pathExists := findConfigs(fileName)
	if pathExists {
		jsonFile, _ := ioutil.ReadFile(pathToFile)
		err := json.Unmarshal(jsonFile, &app)
		if err != nil {
			fmt.Println("FATAL Unmarshaling json config file. Please check the file for sintax errors")
			os.Exit(1)
		}
	} else {
		fmt.Println("Ups! There is no config file!\nConfig file should be in the following paths:\n." + fileName + ".json\n~/.configs/" + fileName + "\n~/." + fileName + " \n/usr/local/etc/" + fileName)
		os.Exit(1)
	}
	return app
}

// Checks in a possivel set of paths if the config file exists
func findConfigs(fileName string) (string, bool) {
	homeDir, _ := homedir.Dir()
	possiblePaths := []string{fileName,
		homeDir + "/.config/" + fileName,
		homeDir + "/.applications.json", "/usr/local/etc/." + fileName}

	for _, possiblePaths := range possiblePaths {
		_, err := os.Stat(possiblePaths)
		pathNotExist := os.IsNotExist(err)
		if !pathNotExist {
			return possiblePaths, true
		}
	}
	return "", false
}

// Extracts from the application configuration
// the given app dependencies and its attributes
func getAppDependencies(appsConfig Applications, dependenciesToVisit []string, visitedDependencies []ApplicationAttributes) []ApplicationAttributes {

	for _, dependencyToVisit := range dependenciesToVisit {
		if sliceContainsString(getSliceOfAppNames(visitedDependencies), dependencyToVisit) {
			continue
		}
		appAttr := getAppAttributes(appsConfig, dependencyToVisit)
		visitedDependencies = append(visitedDependencies, appAttr)
		visitedDependencies = getAppDependencies(appsConfig, appAttr.Dependencies, visitedDependencies)
	}
	return visitedDependencies
}

func getAppAttributes(appsConfig Applications, appName string) ApplicationAttributes {
	for _, appAttr := range appsConfig.ApplicationsMap {
		if appName == appAttr.Name {
			return appAttr
		}
	}
	return ApplicationAttributes{}
}

func sliceContainsString(slice []string, str string) bool {
	for _, element := range slice {
		if str == element {
			return true
		}
	}
	return false
}

// Retrieves a list of the app names from a given []ApplicationAttributes
func getSliceOfAppNames(appsAttributes []ApplicationAttributes) []string {
	appNames := []string{}

	for _, appAttribute := range appsAttributes {
		appNames = append(appNames, appAttribute.Name)
	}
	return appNames
}
