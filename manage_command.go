package main

import (
	"fmt"
)

func unloadCmd(puzzles *[]taquin, args []string) {
	for _, arg := range args {
		var tmpPuzzles []taquin = *puzzles

		for i, puzzle := range *puzzles {
			if puzzle.ID == arg {
				tmpPuzzles = append(tmpPuzzles[:i], tmpPuzzles[i+1:]...)
			}
		}
		*puzzles = tmpPuzzles
	}
}

func showCmd(puzzles []taquin, args []string) int {
	if args == nil {
		for _, puzzle := range puzzles {

			fmt.Printf("%s: %d\n", puzzle.ID, puzzle.size)
		}
		return 0
	}

	for _, arg := range args {
		for _, puzzle := range puzzles {
			if puzzle.ID == arg {
				println(puzzle.size)
				showPuzzle(puzzle)
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
