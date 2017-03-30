package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var timeAgo string
var showDiff bool

var SherlockCmd = &cobra.Command{
	Use:   "sherlock [repo name]",
	Short: "A tool that will help you to quickly identify breaking changes",
	Long: `
# Sherlock

**Code more a worry less about debugging and breaking the application.**

Sherlock is a tool that will quickly search for breaking 
commits in a given app and its dependencies.


## How does it work?
1. Define the dependencies between applications.
2. Call sherlock with the appliaction you want to debug an since when the misbehaviour started
3. Quickly find what might had cause the error

ex ./sherlock application_name --since "1 day"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			sherlock(args[0], timeAgo, showDiff)
			os.Exit(0)
		}
		fmt.Println("Missing argument repo name. Type -h for help")
	},
}

func main() {
	Log.Init(true)
	InitCLI()
}

func InitCLI() {
	SherlockCmd.PersistentFlags().StringVarP(&timeAgo, "since", "t", "1 week", "since when the application is broken (ex: --since \"2 days\")")
	SherlockCmd.PersistentFlags().BoolVarP(&showDiff, "show-diff", "d", false, "show the diff introduced by each commit (ex: --show-diff)")
	SherlockCmd.Execute()
}
