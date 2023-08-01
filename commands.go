package main

import (
	"os"
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
	function    func()
}

func registerCommand(commandName string, description string, args int, function func()) {
	commandList = append(commandList, CommandProperties{
		name:        commandName,
		description: description,
		args:        args,
		function:    function,
	})
}

func initCommands() {
	registerCommand("help", "Show this help menu", 0, help)
	registerCommand("exit", "Exit the application", 0, exit)
}

///////////////////////
// Start of commands //
///////////////////////

func help() {
	for i, props := range commandList {
		print(prefix(0) + color.Green + strings.ToUpper(props.name) + GRAY + " - " + color.Reset + props.description)
		i = i
	}
}

func exit() {
	os.Exit(0)
}
