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
	command     string
	args        []string
	commandList []CommandProperties
)

///////////////////
// Main function //
///////////////////

func main() {
	initCommands()

	print(color.Green + "     _   __ _  __")
	print(color.Green + "    / | / /(_)/ /_ _____ ____ ")
	print(color.Green + "   /  |/ // // __// ___// __ \\" + GRAY + "   Version: " + color.Green + "1.0.0")
	print(color.Green + "  / /|  // // /_ / /   / /_/ /" + GRAY + "   Packages: " + color.Green + "")
	print(color.Green + " /_/ |_//_/ \\__//_/    \\____/ " + GRAY + "   Made by: " + color.Green + "NoahOnFyre")
	print("")
	
	for {
		command, args = commandInput(*scanner)

		for i, props := range commandList {
			if props.name == command {
				props.function()
			}
			i = i
		}
	}
}
