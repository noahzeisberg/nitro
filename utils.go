package main

import (
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func input(msg string) string {
	printR(msg)
	scanner.Scan()
	return scanner.Text()
}

func single(object string, err error) string {
	if err != nil {
		print(prefix(2) + err.Error())
		os.Exit(1)
	}
	return object
}

func commandInput() (string, []string) {
	commandLine := GRAY + "\\\\" + color.Green + "nitro" + GRAY + " ~ " + color.Reset

	userInput := input(commandLine)
	splitted := strings.Split(userInput, " ")
	command := splitted[0]
	args := removeElement(splitted, 0)
	return command, args
}
