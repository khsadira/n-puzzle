package main

func main() {
	var puzzles []taquin

	var puzzle taquin

	puzzle = createPuzzleTest("n-puzzle-1x1", 1)
	puzzle.Size = 1
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-1x1", 1)
	puzzle.Size = 1
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	showPuzzle(puzzle)
	puzzle = createPuzzleTest("n-puzzle-2x2", 2)
	puzzle.Size = 2
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-4x4", 4)
	showPuzzle(puzzle)
	puzzle.Size = 4
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-5x5", 5)
	puzzle.Size = 5
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-5x5-2", 5)
	puzzle.Size = 5
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-5x5-3", 5)
	puzzle.Size = 5
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-3x3", 3)
	puzzle.Size = 3
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-3x3-3", 3)
	puzzle.Size = 3
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-3x3-2", 3)
	puzzle.Size = 3
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-1x1", 1)
	puzzle.Size = 1
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	// puzzle = createPuzzleTest("n-puzzle-15x15", 15)
	// puzzle.Size = 15
	// if isValidTaquin(puzzle) {
	// 	appendPuzzleToPuzzles(&puzzles, puzzle)
	// }

	puzzle = createPuzzleTest("n-puzzle-15x15-2", 15)
	puzzle.Size = 15
	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	puzzle = createPuzzleTest("n-puzzle-test2", 2)
	showPuzzle(puzzle)
	puzzle.Size = 2

	if isValidTaquin(puzzle) {
		appendPuzzleToPuzzles(&puzzles, puzzle)
	}

	// puzzle = createPuzzleTest("n-puzzle-17x17", 17)
	// showPuzzle(puzzle)
	// puzzle.Size = 17
	// if isValidTaquin(puzzle) {
	// 	appendPuzzleToPuzzles(&puzzles, puzzle)
	// }

	gui(&puzzles)

	// for {
	// 	cmd, args := getUserEntry("> ")

	// 	switch cmd {
	// 	case "help":
	// 		helpCmd(args)
	// 	case "show":
	// 		showCmd(puzzles, args)
	// 	case "load":
	// 		loadCmd(&puzzles, args)
	// 	case "unload":
	// 		unloadCmd(&puzzles, args)
	// 	case "gui": // to be deleted
	// 		gui(&puzzles)
	// 	case "start": // to be deleted
	// 		start()
	// 	case "solve":
	// 		println("solve", args)
	// 	case "play":
	// 		playCmd(puzzles, args)
	// 	case "credentials":
	// 		credentialsCmd()
	// 	case "quit":
	// 		return
	// 	default:
	// 		println("n-puzzle: " + cmd + ": command not found\nType `help name' to find out more about the function `name`.")
	// 	}
	// 	println()
	// }
}
