package main

import (
	"bufio"
	"os"

	"github.com/TwiN/go-color"
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
)

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
	})

	print(color.Green + "     _   __ _  __")
	print(color.Green + "    / | / /(_)/ /_ _____ ____ ")
	print(color.Green + "   /  |/ // // __// ___// __ \\" + GRAY + "   Version: " + color.Green + "1.0.0")
	print(color.Green + "  / /|  // // /_ / /   / /_/ /" + GRAY + "   Packages: " + color.Green + "")
	print(color.Green + " /_/ |_//_/ \\__//_/    \\____/ " + GRAY + "   Made by: " + color.Green + "NoahOnFyre")
	print("")

	var command string
	var args []string

	args = args

	for {
		command, args = commandInput()

		for _, props := range commandList {
			if props.name == command {
				props.run()
			}
		}
	}
}
