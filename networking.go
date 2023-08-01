package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getContent(url string) string {
	res, err := http.Get(url)
	if err != nil {
		print(prefix(2) + "HTTP request failed: No response of host.")
		return ""
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		print(prefix(2) + "HTTP request failed: Invalid status.")
		return ""
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		print(prefix(2) + "Reading body failed.")
		return ""
	}

	return string(content)
}

func apiRequest(path string) http.Response {
	json, _ := json.Marshal(requestGet("https://api.github.com" + path))

}

func getRepoContents() {
	apiRequest("/repos/NoahOnFyre/FyUTILS/contents")
}
