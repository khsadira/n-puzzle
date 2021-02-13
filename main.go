package main

import (
	"fmt"
	"strings"
)

type puzzleStruct struct {
	ID     string
	size   int
	puzzle [][]int
}

func showCmd(puzzles []puzzleStruct, args []string) int {
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
				for i := 0; i < puzzle.size; i++ {
					fmt.Printf("%v\n", puzzle.puzzle[i])
				}
			}
		}
		println()
	}

	return 0
}

func createPuzzleTest(ID string, size int) puzzleStruct {
	var puzzle puzzleStruct

	nPuzzle := make([][]int, size)
	puzzle.size = size
	puzzle.ID = ID

	for i := 0; i < size; i++ {
		nPuzzle[i] = make([]int, size)
		for j := 0; j < size; j++ {
			nPuzzle[i][j] = i*i + j
		}
	}

	puzzle.puzzle = nPuzzle

	return puzzle
}

func main() {
	var puzzles []puzzleStruct
	var userEntry string
	var args []string

	puzzles = append(puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-2x2", 2))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-3x3", 3))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-4x4", 4))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-5x5", 5))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-6x6", 6))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-7x7", 7))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-8x8", 8))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-9x9", 9))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-10x10", 10))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-test10", 10))
	puzzles = append(puzzles, createPuzzleTest("n-puzzle-test2", 2))

	for i := 0; i < 1; i++ {
		print("Command: ")
		println()
		userEntry = "show n-puzzle-1x1 n-puzzle-6x6 n-puzzle-test10"
		userEntryArr := strings.Split(userEntry, " ")

		cmd := userEntryArr[0]
		if len(userEntryArr) > 1 {
			args = userEntryArr[1:]
		}

		switch cmd {
		case "help":
			println("help", args)
		case "show":
			showCmd(puzzles, args)
		case "load":
			println("load", args)
		case "solve":
			println("solve", args)
		case "credentials":
			println("lvasseur...\nkhsadira...\n")
		case "quit":
			println("quit reached")
			return
		default:
			println("je n'ai pas compris", cmd)
		}
	}
}
