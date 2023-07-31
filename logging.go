package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)

var RESET = color.Reset
var BLACK = color.Black
var RED = color.Red
var GREEN = color.Green
var YELLOW = color.Yellow
var BLUE = color.Blue
var MAGENTA = color.Purple
var CYAN = color.Cyan
var WHITE = color.White
var GRAY = color.Gray

var BG_RESET = color.Reset
var BG_BLACK = color.BlackBackground
var BG_RED = color.RedBackground
var BG_GREEN = color.GreenBackground
var BG_YELLOW = color.YellowBackground
var BG_BLUE = color.BlueBackground
var BG_MAGENTA = color.PurpleBackground
var BG_CYAN = color.CyanBackground
var BG_WHITE = color.WhiteBackground
var BG_GRAY = color.GrayBackground

func print(msg string) {
	fmt.Println(prefix(0) + msg)
	fmt.Println()
}

func warn(msg string) {
	fmt.Println(prefix(1) + msg)
}

func err(msg string) {
	fmt.Println(prefix(2) + msg)
}

func prefix(level int) string {
	switch level {
	case 0:
		return BG_GREEN + "" + BLACK + " INFO " + BG_RESET + WHITE + " "
	case 1:
		return BG_YELLOW + BLACK + " WARN " + BG_RESET + YELLOW + " "
	case 2:
		return BG_RED + BLACK + " ERROR " + BG_RESET + RED + " "
	}
	return ""
}
