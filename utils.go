package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func input(msg string) string {
	printR(msg)
	scanner.Scan()
	return scanner.Text()
}

func single(object string, err error) string {
	if err != nil {
		print(prefix(2) + err.Error())
		os.Exit(1)
	}
	return object
}

func commandInput() (string, []string) {
	commandLine := GRAY + "\\\\" + color.Green + "nitro" + GRAY + " ~ " + color.Reset

	userInput := input(commandLine)
	splitted := strings.Split(userInput, " ")
	command := splitted[0]
	args := removeElement(splitted, 0)
	return command, args
}

func getManifest(pkg string) Manifest {
	file_content, err := os.ReadFile(package_dir + "\\" + pkg + "\\manifest.npkg")

	if err != nil {
		print(prefix(2) + "Reading manifest data failed! " + err.Error())
		os.Exit(1)
	}

	var manifest Manifest

	json.Unmarshal(file_content, &manifest)

	return manifest
}

func parseRepoName(reponame string) string {
	owner := strings.ToLower(strings.Split(reponame, "/")[0])
	repo := strings.ToLower(strings.Split(reponame, "/")[1])
	return owner + "." + repo
}

func parseLocalRepoName(reponame string) string {
	owner := strings.ToLower(strings.Split(reponame, ".")[0])
	repo := strings.ToLower(strings.Split(reponame, ".")[1])
	return owner + "/" + repo
}
