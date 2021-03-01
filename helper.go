package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func getVoidPosTaquin(taquin [][]uint16, size uint8) (Vector2D, error) {
	var voidpos Vector2D
	var i, j uint8

	for i = 0; i < size; i++ {
		for j = 0; j < size; j++ {
			if taquin[i][j] == 0 {
				voidpos.x = i
				voidpos.y = j
				return voidpos, nil
			}
		}
	}
	return voidpos, errors.New("puzzle not well formatted.")
}

func appendPuzzleToPuzzles(puzzles *[]taquin, puzzleToAdd taquin) {
	for _, puzzle := range *puzzles {
		if puzzle.ID == puzzleToAdd.ID {
			println("n-puzzle: error: Failed to add", puzzleToAdd.ID, "due to a similar ID finded.")
			return
		}
	}

	*puzzles = append(*puzzles, puzzleToAdd)
	println("n-puzze:", puzzleToAdd.ID, "added to puzzles.")
}

func showPuzzle(puzzle taquin) {
	for i := uint8(0); i < puzzle.size; i++ {
		var IDs []string
		for j := uint8(0); j < puzzle.size; j++ {
			IDs = append(IDs, strconv.Itoa(int(puzzle.taquin[i][j])))
		}
		fmt.Println(strings.Join(IDs, " "))
	}
}

func isTaquinSolved(puzzle taquin) bool {
	var taquinArray []uint16 = convertTaquinToArray(puzzle.taquin)
	jmp := 0

	if taquinArray[0] != 0 {
		return false
	}
	for i := 0; i < len(taquinArray)-1; i++ {
		if taquinArray[i]+1 != taquinArray[i+1] {
			jmp++
		}
	}
	if jmp > 1 {
		return false
	}
	return true
}
