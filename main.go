package main

import (
	"github.com/NoahOnFyre/gengine/color"
	"github.com/NoahOnFyre/gengine/convert"
	"github.com/NoahOnFyre/gengine/filesystem"
	"github.com/NoahOnFyre/gengine/logging"
	"github.com/NoahOnFyre/gengine/utils"
	"github.com/google/go-github/github"
	"os"
	"strings"
)

var (
	gitHubClient = github.NewClient(nil)
	version      = "1.1.0"
	userDir, _   = os.UserHomeDir()
	nitroDir     = userDir + "\\.nitro\\"
	pkgDir       = nitroDir + "packages"
	commands     []Command
)

func main() {
	logging.Log("Starting...")
	logging.SetMainColor(color.GreenBg)

	logging.Log("Checking paths...")
	paths := []string{nitroDir, pkgDir}
	for _, path := range paths {
		if !filesystem.Exists(path) {
			err := os.MkdirAll(path, os.ModeDir)
			if err != nil {
				logging.Error("Failed to create Nitro paths.")
				Exit(1)
			}
		}
	}

	packages, err := os.ReadDir(pkgDir)
	if err != nil {
		logging.Error("Failed to load packages in package directory!")
		Exit(1)
	}

	logging.Log("Registering commands...")
	CommandRegistration()

	logging.PrintR("\033[H\033[2J")

	utils.SetTitle("Nitro Package Manager - v" + version + " - " + nitroDir)

	logging.Print(color.Green + "    _   __ _  __")
	logging.Print(color.Green + "   / | / /(_)/ /_ _____ ____   ")
	logging.Print(color.Green + "  /  |/ // // __// ___// __ \\  " + color.Gray + "Made by: " + color.Green + "NoahOnFyre")
	logging.Print(color.Green + " / /|  // // /_ / /   / /_/ /  " + color.Gray + "Version: " + color.Green + version)
	logging.Print(color.Green + "/_/ |_//_/ \\__//_/    \\____/   " + color.Gray + "Packages: " + color.Green + convert.FormatInt(len(packages)))
	logging.Print(color.Green + "")

	for {
		raw := logging.Input(color.Gray + "\\\\" + color.Green + " nitro " + color.Gray + "~" + color.Reset + " ")
		logging.Print()

		split := strings.Split(raw, " ")
		command := split[0]
		args := utils.RemoveElement(split, 0)

		for _, props := range commands {
			if props.Name == command {
				if len(args) == props.Args.Count {
					props.Run(args)
					logging.Print()
				} else {
					logging.Error("Unexpected arguments!")
					logging.Print()
				}
			}
		}
	}
}

func Exit(code int) {
	os.Exit(code)
}
