package main

import (
	"io/fs"
	"os"
)

func checkPaths(paths []string) {
	for _, path := range paths {
		if !exists(path) {
			os.Mkdir(path, fs.ModeDir)
		}
	}
}
