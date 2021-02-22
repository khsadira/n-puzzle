package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	return true
}
