package main

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
