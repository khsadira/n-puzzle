package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
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

func convertFileToPuzzle(file io.Reader) (taquin, error) {
	var puzzle taquin
	var err error

	scanner := bufio.NewScanner(file)
	var i uint8

	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			var size64 uint64

			size64, err = strconv.ParseUint(line, 10, 64)

			if err != nil || size64 <= 0 {
				return puzzle, errors.New("puzzle not well formatted.")
			}

			puzzle.size = uint8(size64)
			puzzle.taquin = make([][]uint16, puzzle.size)
		} else if i > puzzle.size {
			return puzzle, errors.New("puzzle not well formatted.")
		} else {
			puzzle.taquin[i-1], err = strToMapInt(line, puzzle.size)

			if err != nil {
				return puzzle, errors.New("puzzle not well formatted.")
			}
		}
		i++
	}

	if i != puzzle.size+1 {
		return puzzle, errors.New("puzzle not well formatted.")
	}

	return puzzle, nil
}

func loadFilePuzzleToPuzzles(puzzles *[]taquin, file *os.File, arg string) {
	var puzzle taquin

	puzzle, err := convertFileToPuzzle(file)
	if err != nil {
		println("n-puzzle: load:", arg, err.Error())
		return
	}

	puzzle.voidpos, err = getVoidPosTaquin(puzzle.taquin, puzzle.size)

	if err != nil {
		println("n-puzzle: load:", arg, err.Error())
		return
	}

	puzzle.ID = filepath.Base(arg)

	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(puzzles, puzzle)
	}
}

func loadCmd(puzzles *[]taquin, args []string) {
	var file *os.File
	var err error

	for _, arg := range args {
		file, err = os.Open(arg)

		if err != nil {
			println("n-puzzle: load: error:", err.Error())
			continue
		}

		defer file.Close()
		loadFilePuzzleToPuzzles(puzzles, file, arg)
	}
}
