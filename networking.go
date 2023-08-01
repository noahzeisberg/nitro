package main

import (
	"net/http"
)

func apiRequest(path string) {
	result, err := http.Get("https://api.github.com" + path)
	if err != nil {
		error("GitHub API request failed!")
		os.Exit(1)
	}
	return result
}

func getRepoContents() {
	 apiRequest("/repos/NoahOnFyre/FyUTILS/contents").Json()
}