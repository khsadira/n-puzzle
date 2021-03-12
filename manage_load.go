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

			puzzle.Size = uint8(size64)
			puzzle.Taquin = make([][]uint16, puzzle.Size)
		} else if i > puzzle.Size {
			return puzzle, errors.New("puzzle not well formatted.")
		} else {
			puzzle.Taquin[i-1], err = strToMapInt(line, puzzle.Size)

			if err != nil {
				return puzzle, errors.New("puzzle not well formatted.")
			}
		}
		i++
	}

	if i != puzzle.Size+1 {
		return puzzle, errors.New("puzzle not well formatted.")
	}

	return puzzle, nil
}

func loadFileDataToGlobalData(file *os.File, arg string) {
	var data metaTaquin
	var err error

	data.TaquinStruct, err = convertFileToPuzzle(file)
	if err != nil {
		println("n-puzzle: load:", arg, err.Error())
		return
	}

	data.TaquinStruct.Voidpos, err = getVoidPosTaquin(data.TaquinStruct.Taquin, data.TaquinStruct.Size)

	if err != nil {
		println("n-puzzle: load:", arg, err.Error())
		return
	}

	data.ID = filepath.Base(arg)

	if isValidTaquin(data.ID, data.TaquinStruct) {
		appendDataToGlobalData(data)
	}
}

func loadCmd(args []string) {
	var file *os.File
	var err error

	for _, arg := range args {
		file, err = os.Open(arg)

		if err != nil {
			println("n-puzzle: load: error:", err.Error())
			continue
		}

		defer file.Close()
		loadFileDataToGlobalData(file, arg)
	}
}
