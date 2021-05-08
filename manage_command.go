package main

import (
	"fmt"
	"strconv"
)

func generateCmd(args []string) {
	if len(args) != 2 {
		println("generate: usage: [name] [size]")
		return
	}
	size, _ := strconv.Atoi(args[1])

	if size >= 2 {
		var data metaTaquin
		var err error

		data.TaquinStruct = generate_taquin(uint8(size))
		mix_taquin(&data.TaquinStruct)
		data.TaquinStruct.Voidpos, err = getVoidPosTaquin(data.TaquinStruct.Taquin, data.TaquinStruct.Size)

		if err != nil {
			println("n-puzzle: generate:", err.Error())
			return
		}

		data.ID = args[0]
		appendDataToGlobalData(data)
	} else {
		println("n-puzzle: generate: Your taquin size must be superior or equal to 2.")
	}
}

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
				println()
			}
		}
	}

	return 0
}

func helpCmd(args []string) int {
	if args == nil {
		print("help\ngenerate\nshow\nload\nunload\nenv\nset\nsolve\nplay\ngui\ncredentials\nquit\n")
		return 0
	}

	for _, arg := range args {
		switch arg {
		case "generate":
			println("generate: usage: generate [name] [size]")
		case "help":
			println("help: usage: help [] - Print the list of each usable commands by the programs\nhelp [cmd] - Print the usage of the command")
		case "show":
			println("show: usage: show [] - print the puzzle list | show [puzzleID] - Print the puzzle")
		case "load":
			println("load: usage: load [args] - Load a puzzle from a file")
		case "unload":
			println("unload: usage: unload [args] - Unload a puzzle from the data")
		case "solve":
			println("solve: usage: solve [args] - Solve the taquin")
		case "env":
			println("env: usage: env [] - Show current environment")
		case "set":
			println("set: usage: set heur=X algo=X - Set env variable")
		case "play":
			println("play: usage: play [puzzleID] - Start the n-puzzle game with the selected puzzle ID")
		case "gui":
			println("gui: usage: gui [] - Start the GU interface")
		case "credentials":
			println("credentials: usage: credentials [] - Show the coredentials")
		case "quit":
			println("quit: usage: quit [] - Leave the program")
		default:
			println(arg + ": command not found\nType help to see every commands.")
		}
	}
	return 0
}

func credentialsCmd() {
	println("lvasseur...\nkhsadira...\n")
}
