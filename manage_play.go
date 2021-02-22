package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func startGame(puzzle taquin) {
	println("Welcome in the play room ! Rules are easy, press arrow keys to move the puzzle, esc to quit.\nYou are using the next puzzle.")

	// startTime := time.Now()
	plays := 0
	taquinSolve := false

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for !taquinSolve {
		showPuzzle(puzzle)
		println()

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowLeft:
			move_left(&puzzle)
		case keyboard.KeyArrowRight:
			move_right(&puzzle)
		case keyboard.KeyArrowUp:
			move_up(&puzzle)
		case keyboard.KeyArrowDown:
			move_down(&puzzle)
		case keyboard.KeyEsc:
			taquinSolve = true
		default:
			fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		}

		plays++
	}

	// endTime := time.Now()
}

func playCmd(puzzles []taquin, args []string) {
	if len(args) > 0 {
		for _, puzzle := range puzzles {
			if puzzle.ID == args[0] {
				startGame(puzzle)
			}
		}
	}
}
