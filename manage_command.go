package main

import (
	"fmt"
)

func unloadCmd(args []string) {
	for _, arg := range args {
		removeData(arg)
	}
}

func showCmd(args []string) int {
	if args == nil {
		for _, data := range globalData {
			fmt.Printf("%s: %d\n", data.ID, data.TaquinStruct.Size)
		}
		return 0
	}

	for _, arg := range args {
		for _, data := range globalData {
			if data.ID == arg {
				println(data.TaquinStruct.Size)
				showPuzzle(data.TaquinStruct)
			}
		}
	}

	return 0
}

func helpCmd(args []string) int {
	if args == nil {
		print("help\ngenerate\nshow\nload\nsolve\ncredentials\nquit\n")
		return 0
	}

	for _, arg := range args {
		switch arg {
		case "help":
			println("help: usage: help without arguments will print you the list of each usable commands by the programs. help with arguments will print you the usage of the command.")
		case "show":
			println("show: usage:")
		case "load":
			println("load: usage:")
		case "solve":
			println("solve: usage:")
		case "set":
			println("set: usage:")
		case "gui":
			println("gui: usage:")
		case "credentials":
			println("credentials: usage:")
		case "quit":
			println("quit: usage:")
		default:
			println(arg + ": command not found\nType help to see every commands.")
		}
	}
	return 0
}

func credentialsCmd() {
	println("lvasseur...\nkhsadira...\n")
}
