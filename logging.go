package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)

func print(msg any) {
	fmt.Println(msg)
}

func printR(msg any) {
	fmt.Print(msg)
}

func prefix(level int) string {
	switch level {
	case 0:
		return color.GreenBackground + "" + color.Black + " INFO " + color.Reset + color.Reset + " "
	case 1:
		return color.YellowBackground + color.Black + " WARN " + color.Reset + color.Yellow + " "
	case 2:
		return color.RedBackground + color.Black + " ERROR " + color.Reset + color.Red + " "
	}
	return ""
}
