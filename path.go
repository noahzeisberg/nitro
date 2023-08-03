package main

import (
	"io/fs"
	"os"
)

func checkPaths(paths []string) int {
	paths_fixed := 0
	for _, path := range paths {
		if !exists(path) {
			os.Mkdir(path, fs.ModeDir)
			paths_fixed += 1
		}
	}
	return paths_fixed
}
