package main

import (
	"context"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/google/go-github/github"
)

var (
	gitHubClient = github.NewClient(nil)
)

func fetchRepo(reponame string) {
	owner := strings.Split(reponame, "/")[0]
	repo := strings.Split(reponame, "/")[1]
	pkg_name := package_dir + "\\" + owner + "-" + repo + "\\"
	if exists(pkg_name) {
		print(prefix(2) + "The package is already installed!")
		return
	}
	os.Mkdir(pkg_name, fs.ModeDir)
	_, directoryContent, response, err := gitHubClient.Repositories.GetContents(context.Background(), owner, repo, "", &github.RepositoryContentGetOptions{})

	if err != nil {
		print(prefix(2) + "Repository not found!")
		return
	}

	for _, contentProperties := range directoryContent {
		if *contentProperties.Type != "file" {
			continue
		}
		file, _ := os.Create(pkg_name + *contentProperties.Name)

		print(prefix(0) + "Fetching " + color.Green + *contentProperties.Name + color.Reset + "... " + GRAY + "(" + strconv.Itoa(contentProperties.GetSize()) + " Bytes" + ")" + GRAY + " - " + color.Green + response.Status)
		bytes_written, err := file.Write([]byte(*contentProperties.Content))
		if err != nil {
			print(prefix(2) + "Error while writing file: " + err.Error())
			return
		}
		print(prefix(0) + "Wrote " + strconv.Itoa(bytes_written) + " bytes to " + file.Name())
		file.Close()
	}
}
