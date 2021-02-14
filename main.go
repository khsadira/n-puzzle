package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type puzzleStruct struct {
	ID     string
	size   uint8
	puzzle [][]uint8
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
				for i := uint8(0); i < puzzle.size; i++ {
					fmt.Printf("%v\n", puzzle.puzzle[i])
				}
			}
		}
	}

	return 0
}

func createPuzzleTest(ID string, size uint8) puzzleStruct {
	var puzzle puzzleStruct

	nPuzzle := make([][]uint8, size)
	puzzle.size = size
	puzzle.ID = ID

	for i := uint8(0); i < size; i++ {
		nPuzzle[i] = make([]uint8, size)
		for j := uint8(0); j < size; j++ {
			nPuzzle[i][j] = i + j
		}
	}

	puzzle.puzzle = nPuzzle

	return puzzle
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

func strToMapInt(line string, size uint8) ([]uint8, error) {
	var err error = nil

	strs := strings.Split(line, " ")
	strlens := len(strs)

	if uint8(strlens) != size {
		return nil, errors.New("puzzle not well formatted.")
	}
	ary := make([]uint8, strlens)

	for i := range ary {
		var value64 uint64
		value64, err = strconv.ParseUint(strs[i], 10, 64)
		ary[i] = uint8(value64)

		if err != nil {
			return nil, err
		}
	}

	return ary, nil
}

func convertFileToPuzzle(file *os.File) (puzzleStruct, error) {
	var puzzle puzzleStruct
	var err error

	scanner := bufio.NewScanner(file)
	var i int

	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			var size64 uint64

			size64, err = strconv.ParseUint(line, 10, 64)

			if err != nil {
				return puzzle, errors.New("puzzle not well formatted.")
			}

			puzzle.size = uint8(size64)
			puzzle.puzzle = make([][]uint8, puzzle.size)
		} else if uint8(i) > puzzle.size {
			return puzzle, errors.New("puzzle not well formatted.")
		} else {
			puzzle.puzzle[i-1], err = strToMapInt(line, puzzle.size)

			if err != nil {
				return puzzle, errors.New("puzzle not well formatted.")
			}
		}
		i++
	}

	return puzzle, nil
}

func loadCmd(puzzles *[]puzzleStruct, args []string) {
	var puzzle puzzleStruct
	var file *os.File
	var err error

	for _, arg := range args {
		file, err = os.Open(arg)

		if err != nil {
			println("n-puzzle: load: error:", err.Error())
			continue
		}

		defer file.Close()

		puzzle, err = convertFileToPuzzle(file)
		if err != nil {
			println("n-puzzle: load:", arg, err.Error())
			continue
		}
		puzzle.ID = arg

		appendPuzzleToPuzzles(puzzles, puzzle)
	}
}

func showPrompt() {
	print("> ")
}

func getUserEntry() (string, []string) {
	var args []string

	showPrompt()
	in := bufio.NewReader(os.Stdin)
	userEntry, err := in.ReadString('\n')

	if err != nil {
		log.Fatal("n-puzzle: error:", err.Error())
	}

	userEntryArr := strings.Split(strings.TrimRight(userEntry, "\r\n"), " ")
	cmd := userEntryArr[0]

	if len(userEntryArr) > 1 {
		args = userEntryArr[1:]
	}

	return cmd, args
}

func appendPuzzleToPuzzles(puzzles *[]puzzleStruct, puzzleToAdd puzzleStruct) {
	for _, puzzle := range *puzzles {
		if puzzle.ID == puzzleToAdd.ID {
			println("n-puzzle: error: Failed to add", puzzleToAdd.ID, "due to a similar ID finded.")
			return
		}
	}

	*puzzles = append(*puzzles, puzzleToAdd)
	println("n-puzze:", puzzleToAdd.ID, "added to puzzles.")
}

func main() {
	var puzzles []puzzleStruct

	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-2x2", 2))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-4x4", 4))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-5x5", 5))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-3x3", 3))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-10x10", 10))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-test2", 2))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))

	for {
		cmd, args := getUserEntry()

		switch cmd {
		case "help":
			helpCmd(args)
		case "show":
			showCmd(puzzles, args)
		case "load":
			loadCmd(&puzzles, args)
		case "solve":
			println("solve", args)
		case "play":
			println("play", args)
		case "credentials":
			credentialsCmd()
		case "quit":
			return
		default:
			println("n-puzzle: " + cmd + ": command not found\nType `help name' to find out more about the function `name`.")
		}
	}
}
