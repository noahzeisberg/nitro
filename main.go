package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/TwiN/go-color"
	title "github.com/lxi1400/GoTitle"
)

//////////////////////
// Global variables //
//////////////////////

var (
	scanner     = bufio.NewScanner(os.Stdin)
	commandList []CommandProperties

	user_dir    string = single(os.UserHomeDir())
	nitro_dir   string = user_dir + "\\.nitro"
	core_dir    string = nitro_dir + "\\core"
	package_dir string = nitro_dir + "\\packages"
	temp_dir    string = nitro_dir + "\\temp"
	plugin_dir  string = core_dir + "\\plugins"

	command string
	args    []string
)

func menu() {
	dir_content, err := os.ReadDir(package_dir)

	if err != nil {
		print(prefix(2) + "The package directory couldn't be read.")
	}

	printR("\033[H\033[2J")
	title.SetTitle("Nitro Package Manager - " + nitro_dir)
	print(color.Green + "     _   __ _  __")
	print(color.Green + "    / | / /(_)/ /_ _____ ____ ")
	print(color.Green + "   /  |/ // // __// ___// __ \\" + GRAY + "   Version: " + color.Green + "1.0.0")
	print(color.Green + "  / /|  // // /_ / /   / /_/ /" + GRAY + "   Packages: " + color.Green + strconv.Itoa(len(dir_content)))
	print(color.Green + " /_/ |_//_/ \\__//_/    \\____/ " + GRAY + "   Made by: " + color.Green + "NoahOnFyre")
}

///////////////////
// Main function //
///////////////////

func main() {
	initCommands()

	checkPaths([]string{
		nitro_dir,
		core_dir,
		package_dir,
		temp_dir,
		plugin_dir,
	})

	menu()
	print("")

	args = args

	for {
		command, args = commandInput()
		print("")

		for _, props := range commandList {
			if props.name == command {
				if len(args) == props.args {
					props.run()
					print("")
				} else {
					print(prefix(2) + "Unexpected arguments!")
					print("")
				}
			}
		}
	}
}
