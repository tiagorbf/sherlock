package main

import "fmt"
import "errors"

func sherlock(appName, timeAgo string, showDiff bool) {
	timeAgo += " ago"

	// Get a structure of the configuration file
	app := initConfig()

	// Get app dependencies
	appDependencies := getAppDependencies(app, []string{appName}, []ApplicationAttributes{})

	if len(appDependencies) == 0 {
		err := errors.New("Can not find " + appName + " in config file")
		Log.WriteLine("FATAL", "Getting dependencies", err)
	}

	for _, appDependency := range appDependencies {
		getRepo(appDependency.Path, appDependency.CloneUrl)

		fmt.Println("===============================================")
		fmt.Println("            " + appDependency.Name)
		fmt.Println("===============================================")
		fmt.Printf("%s\n\n", getCommitsSince(appDependency.Path, timeAgo, showDiff))
	}
}
