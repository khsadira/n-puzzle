package main

func main() {
	var puzzles []taquin

	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-2x2", 2))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-4x4", 4))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-5x5", 5))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-3x3", 3))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-15x15", 15))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-16x16", 16))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-17x17", 17))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-test2", 2))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))
	appendPuzzleToPuzzles(&puzzles, createPuzzleTest("n-puzzle-1x1", 1))

	for {
		cmd, args := getUserEntry("> ")

		switch cmd {
		case "help":
			helpCmd(args)
		case "show":
			showCmd(puzzles, args)
		case "load":
			loadCmd(&puzzles, args)
		case "unload":
			unloadCmd(&puzzles, args)
		case "start":
			start()
		case "solve":
			println("solve", args)
		case "play":
			playCmd(&puzzles, args)
		case "credentials":
			credentialsCmd()
		case "quit":
			return
		default:
			println("n-puzzle: " + cmd + ": command not found\nType `help name' to find out more about the function `name`.")
		}
		println()
	}
}
