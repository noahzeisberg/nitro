package main

import (
	"context"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/TwiN/go-color"
	"github.com/google/go-github/github"
)

var (
	gitHubClient = github.NewClient(nil)
)

type Manifest struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Type     string `json:"type"`
	Build    string `json:"build"`
	Main     string `json:"main"`
}

func fetchRepo(reponame string) {
	owner := strings.ToLower(strings.Split(reponame, "/")[0])
	repo := strings.ToLower(strings.Split(reponame, "/")[1])
	pkg_name := owner + "." + repo
	pkg_dir := package_dir + "\\" + owner + "." + repo + "\\"
	if exists(pkg_dir) {
		print(prefix(2) + "The package is already installed!")
		return
	}
	_, directoryContent, response, err := gitHubClient.Repositories.GetContents(context.Background(), owner, repo, "", &github.RepositoryContentGetOptions{})

	if err != nil {
		print(prefix(2) + "Repository not found!")
		return
	}

	os.Mkdir(pkg_dir, fs.ModeDir)

	var wg sync.WaitGroup
	wg.Add(len(directoryContent))

	for _, contentProps := range directoryContent {
		go func(props *github.RepositoryContent) {
			if *props.Type != "file" {
				wg.Done()
				return
			}
			file, _ := os.Create(pkg_dir + *props.Name)

			print(prefix(0) + "Fetching " + color.Green + *props.Name + color.Reset + "... " + GRAY + "(" + strconv.Itoa(props.GetSize()) + " Bytes" + ")" + GRAY + " - " + color.Green + response.Status)

			res, err := http.Get(*props.DownloadURL)
			if err != nil {
				print(prefix(2) + "GET request failed: " + err.Error())
				os.Exit(1)
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)

			if err != nil {
				print(prefix(2) + "Reading body failed: " + err.Error())
				os.Exit(1)
			}

			bytes_written, err := file.Write(body)
			if err != nil {
				print(prefix(2) + "Error while writing file: " + err.Error())
				os.Exit(1)
			}
			print(prefix(0) + "Wrote " + color.Green + strconv.Itoa(bytes_written) + color.Reset + " bytes to " + GRAY + file.Name())
			file.Close()
			wg.Done()
		}(contentProps)
	}
	wg.Wait()

	package_content, err := os.ReadDir(pkg_dir)

	if err != nil {
		print(prefix(2) + "Reading directory failed! " + err.Error())
		os.Exit(1)
	}

	contains_manifest := false

	for _, file := range package_content {
		if file.Name() == "manifest.npkg" {
			contains_manifest = true
		}
	}

	if contains_manifest {
		getManifest(pkg_name)
	}
}
