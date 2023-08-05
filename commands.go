package main

import (
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
)

///////////////////////
// Command utilities //
///////////////////////

type CommandProperties struct {
	name        string
	description string
	args        int
	run         func()
}

func registerCommand(commandName string, description string, args int, run func()) {
	commandList = append(commandList, CommandProperties{
		name:        commandName,
		description: description,
		args:        args,
		run:         run,
	})
}

func initCommands() {
	registerCommand("get", "Fetch all files in the specified repository.", 1, get)
	registerCommand("check", "Check things like your rate limit, internet connection, etc.", 1, check)
	registerCommand("remove", "Remove the data of a fetched repository.", 1, remove)
	registerCommand("list", "List all packages installed.", 0, list)
	registerCommand("dir", "Open the Nitro path.", 0, dir)
	registerCommand("help", "Show this help menu.", 0, help)
	registerCommand("exit", "Exit the application.", 0, exit)
	registerCommand("clear", "Clear the terminal screen.", 0, clear)
}

///////////////////////
// Start of commands //
///////////////////////

func help() {
	for _, props := range commandList {
		print(prefix(0) + color.Green + strings.ToUpper(props.name) + GRAY + " - " + color.Reset + props.description)
	}
}

func exit() {
	os.Exit(0)
}

func get() {
	repo := args[0]
	if !strings.Contains(repo, "/") {
		repo := "noahonfyre/" + repo
		fetchRepo(repo)
	} else {
		fetchRepo(repo)
	}
}

func list() {
	dir_content, err := os.ReadDir(package_dir)

	if err != nil {
		print(prefix(2) + "The package directory couldn't be read.")
	}

	for _, pkg := range dir_content {
		if pkg.IsDir() {
			pkg_info, err := pkg.Info()

			if err != nil {
				print(prefix(2) + "The file information couldn't be read.")
			}
			print(prefix(0) + color.Green + strings.Replace(pkg.Name(), ".", "/", -1) + GRAY + " - " + color.Reset + strconv.Itoa(int(pkg_info.Size())) + " Bytes" + GRAY + " - " + color.Reset + pkg_info.ModTime().Format("15:04:05 - 02.01.2006"))
		}
	}
}

func remove() {
	repo := args[0]
	if !strings.Contains(repo, "/") {
		repo := "NoahOnFyre/" + repo
		os.RemoveAll(package_dir + "\\" + strings.Replace(repo, "/", ".", -1))
	} else {
		os.RemoveAll(package_dir + "\\" + strings.Replace(repo, "/", ".", -1))
	}
	print(prefix(0) + "Package successfully removed!")
}

func check() {
	action := args[0]
	rateLimit, response, _ := gitHubClient.RateLimits(context.Background())

	switch action {
	case "connection":
		print(prefix(0) + "GitHub API returned: " + color.Green + response.Status + color.Reset + " over " + GRAY + response.Proto)

	case "ratelimit":
		print(prefix(0) + "Your current ratelimit: " + color.Green + strconv.Itoa(rateLimit.Core.Remaining) + "/" + strconv.Itoa(rateLimit.Core.Limit))

	case "installation":
		errors_found := 0
		print(prefix(0) + "Checking your Nitro installation... This may take a while.")
		print(prefix(0) + "Checking directories...")
		errors_found += checkPaths([]string{
			nitro_dir,
			core_dir,
			package_dir,
			temp_dir,
			plugin_dir,
		})
		print(prefix(0) + "Checking configuration files...")
		print(prefix(0) + "Checking Go installation...")
		if exists(user_dir + "\\go") {
			print(prefix(0) + "Go is installed!")
		} else {
			print(prefix(2) + "Go is not installed!")
			errors_found += 1
		}
		if errors_found == 0 {
			print(prefix(0) + "Done! " + color.Green + strconv.Itoa(errors_found) + color.Reset + " errors found and fixed!")
		} else {
			print(prefix(1) + "Done! " + color.Green + strconv.Itoa(errors_found) + color.Reset + " errors found and fixed!")
		}

	default:
		print(prefix(2) + "Not a recognized field!")
	}
}

func clear() {
	menu()
}

func dir() {
	exec.Command("start", "explorer.exe " + nitro_dir).Run()
}