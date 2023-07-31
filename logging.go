package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)

func print(msg string) {
	fmt.Println(msg)
}

func printR(msg string) {
	fmt.Print(msg)
}

func prefix(level int) string {
	switch level {
	case 0:
		return color.GreenBackground + "" + color.Black + " INFO " + color.Reset + color.White + " "
	case 1:
		return color.YellowBackground + color.Black + " WARN " + color.Reset + color.Yellow + " "
	case 2:
		return color.RedBackground + color.Black + " ERROR " + color.Reset + color.Red + " "
	}
	return ""
}
