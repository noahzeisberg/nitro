package main

import (
	"bufio"
	"strings"

	"github.com/TwiN/go-color"
)

func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func input(msg string, scanner bufio.Scanner) string {
	printR(msg)
	scanner.Scan()
	return scanner.Text()
}

func commandInput(scanner bufio.Scanner) (string, []string) {
	commandLine := GRAY + "\\\\" + color.Green + "nitro" + GRAY + " ~ " + color.Reset

	userInput := input(commandLine, scanner)
	splitted := strings.Split(userInput, " ")
	command := splitted[0]
	args := removeElement(splitted, 0)
	return command, args
}
