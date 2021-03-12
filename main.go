package main

func main() {
	var puzzle taquin
	var meta metaTaquin

	puzzle, meta = createPuzzleTest("n-puzzle-1x1", 1)
	puzzle.Size = 1
	if isValidTaquin("n-puzzle-1x1", puzzle) {
		appendDataToGlobalData(meta)
	}

	puzzle, meta = createPuzzleTest("n-puzzle-1x1", 1)
	puzzle.Size = 1
	if isValidTaquin("n-puzzle-1x1", puzzle) {
		appendDataToGlobalData(meta)
	}

	showPuzzle(puzzle)
	puzzle, meta = createPuzzleTest("n-puzzle-2x2", 2)
	puzzle.Size = 2
	if isValidTaquin("n-puzzle-2x2", puzzle) {
		appendDataToGlobalData(meta)
	}

	puzzle, meta = createPuzzleTest("n-puzzle-5x5", 5)
	puzzle.Size = 5
	if isValidTaquin("n-puzzle-5x5", puzzle) {
		appendDataToGlobalData(meta)
	}

	puzzle, meta = createPuzzleTest("n-puzzle-3x3", 3)
	puzzle.Size = 3
	if isValidTaquin("n-puzzle-3x3", puzzle) {
		appendDataToGlobalData(meta)
	}

	for {
		cmd, args := getUserEntry("> ")

		switch cmd {
		case "help":
			helpCmd(args)
		case "env":
			println("heuristic:", heur)
			println("algorithm:", algo)
		case "show":
			showCmd(args)
		case "load":
			loadCmd(args)
		case "unload":
			unloadCmd(args)
		case "set":
			setCmd(args)
		case "gui": // to be reworkd
			gui()
		case "start": // to be deleted
			start()
		case "solve":
			for _, arg := range args {
				for _, data := range globalData {
					if data.ID == arg {
						cpy := createPuzzleCopy(data.TaquinStruct)
						algorithm[algo](&cpy)
					}
				}
			}
		case "play":
			playCmd(args)
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
