package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func strToMapInt(line string, size uint8) ([]uint16, error) {
	var err error = nil

	strs := strings.Split(line, " ")
	strlens := len(strs)

	if uint8(strlens) != size {
		return nil, errors.New("puzzle not well formatted.")
	}
	ary := make([]uint16, strlens)

	for i := range ary {
		var value64 uint64
		value64, err = strconv.ParseUint(strs[i], 10, 64)
		ary[i] = uint16(value64)

		if err != nil {
			return nil, err
		}
	}

	return ary, nil
}

func convertFileToPuzzle(file *os.File) (taquin, error) {
	var puzzle taquin
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
			puzzle.taquin = make([][]uint16, puzzle.size)
		} else if uint8(i) > puzzle.size {
			return puzzle, errors.New("puzzle not well formatted.")
		} else {
			puzzle.taquin[i-1], err = strToMapInt(line, puzzle.size)

			if err != nil {
				return puzzle, errors.New("puzzle not well formatted.")
			}
		}
		i++
	}

	return puzzle, nil
}

func loadCmd(puzzles *[]taquin, args []string) {
	var puzzle taquin
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

		if checkTaquin(puzzle) {
			appendPuzzleToPuzzles(puzzles, puzzle)
		}
	}
}

func unloadCmd(puzzles []taquin, args []string) []taquin {
	for _, arg := range args {
		for i, puzzle := range puzzles {
			if puzzle.ID == arg {
				puzzles = append(puzzles[:i], puzzles[i+1:]...)
			}
		}
	}
	return puzzles
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
				for i := uint8(0); i < puzzle.size; i++ {
					var IDs []string
					for j := uint8(0); j < puzzle.size; j++ {
						IDs = append(IDs, strconv.Itoa(int(puzzle.taquin[i][j])))
					}
					fmt.Println(strings.Join(IDs, " "))
				}
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
