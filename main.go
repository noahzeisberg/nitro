package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	print(color.Green + "     _   __ _  __")
	print(color.Green + "    / | / /(_)/ /_ _____ ____ ")
	print(color.Green + "   /  |/ // // __// ___// __ \\" + color.Gray + "   Version: " + color.Green + "1.0.0")
	print(color.Green + "  / /|  // // /_ / /   / /_/ /" + color.Gray + "   Packages: " + color.Green + "")
	print(color.Green + " /_/ |_//_/ \\__//_/    \\____/ " + color.Gray + "   Made by: " + color.Green + "NoahOnFyre")
	print("")
	for {
		printR(color.Gray + "\\\\" + color.Green + "nitro" + color.Gray + " ~ ")
		scanner.Scan()

		userInput := strings.Split(scanner.Text(), " ")
		command := userInput[0]
		// args := removeElement(userInput, 0)

		switch command {
		case "exit":
			os.Exit(0)
		}
	}
}
