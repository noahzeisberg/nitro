package main

import (
	"context"
	"io/fs"
	"os"
	"strings"

	"github.com/google/go-github/github"
)

var (
	client = github.NewClient(nil)
)

func fetchRepo(reponame string) {
	owner := strings.Split(reponame, "/")[0]
	repo := strings.Split(reponame, "/")[1]
	pkg_name := package_dir + "\\" + owner + "-" + repo + "\\"
	os.Mkdir(pkg_name, fs.ModeDir)
	_, directoryContent, _, err := client.Repositories.GetContents(context.Background(), owner, repo, "", &github.RepositoryContentGetOptions{})

	for _, content := range directoryContent {
		os.Create(pkg_name + *content.Name)
	}

	if err != nil {
		print(prefix(2) + "Repository not found!")
		return
	}
}
