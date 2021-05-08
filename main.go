package main

func main() {
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
		case "generate":
			generateCmd(args)
		case "load":
			loadCmd(args)
		case "unload":
			unloadCmd(args)
		case "set":
			setCmd(args)
		case "gui": // to be reworkd
			gui()
		case "solve":
			for _, arg := range args {
				for _, data := range globalData {
					if data.ID == arg {
						println("Start solve:", data.ID)
						cpy := createPuzzleCopy(data.TaquinStruct)
						algorithm[algo](&cpy)
						println()
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
