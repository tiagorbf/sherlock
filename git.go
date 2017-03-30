package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func getRepo(pathToRepo, cloneUrl string) {
	if fileExists(pathToRepo) {
		gitPull(pathToRepo)
	} else {
		gitClone(cloneUrl, pathToRepo)
	}
}

func gitClone(repo, pathToRepo string) {
	Log.WriteLine("LOG", "Cloning repository: "+repo, nil)
	_, err := exec.Command("/usr/bin/git", "clone", repo, pathToRepo).Output()
	Log.WriteLine("FATAL", "Cloning repository: "+repo, err)
	Log.WriteLine("LOG", "Done cloning release_queries repository", nil)
}

func gitPull(pathToRepo string) {
	currentDir := getCurrentDir()
	cdToDir(pathToRepo)
	Log.WriteLine("LOG", "Pulling "+pathToRepo+" repository", nil)
	_, err := exec.Command("/usr/bin/git", "pull").Output()
	Log.WriteLine("FATAL", "git pull", err)
	cdToDir(currentDir)
}

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	Log.WriteLine("FATAL", "Getting absolute path to binary", err)
	return dir
}

func getCommitsSince(pathToRepo, timeAgo string, showDiff bool) string {
	currentDir := getCurrentDir()
	cdToDir(pathToRepo)

	result, err := exec.Command("/usr/bin/git", getGitLogArgs(timeAgo, showDiff)...).Output()
	Log.WriteLine("WARNING", "Fail to execute git log", err)
	cdToDir(currentDir)
	return string(result[:])
}

func getGitLogArgs(timeAgo string, showDiff bool) []string {
	gitLogArgs := []string{
		"log",
		"--pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset'",
		"--since='" + timeAgo + "'",
		"--color=always",
	}
	if showDiff {
		gitLogArgs = append(gitLogArgs, "-p")
	}
	return gitLogArgs
}

func fileExists(pathToFile string) bool {
	_, err := os.Stat(pathToFile)
	pathNotExist := os.IsNotExist(err)
	if pathNotExist {
		return false
	}
	return true
}

func cdToDir(pathToDir string) {
	err := os.Chdir(pathToDir)
	Log.WriteLine("FATAL", "Open dir release_queries", err)
}
