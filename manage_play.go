package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func startGame(ID string, puzzle taquin) {
	println("Welcome in the play room ! Rules are easy, press ARROW keys to move the puzzle, ENTER to valide your puzzle or ESC to quit.\nYou are using the next puzzle.")

	startTime := time.Now().Unix()
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

		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowLeft:
			move_left(&puzzle)
			plays++
		case keyboard.KeyArrowRight:
			move_right(&puzzle)
			plays++
		case keyboard.KeyArrowUp:
			move_up(&puzzle)
			plays++
		case keyboard.KeyArrowDown:
			move_down(&puzzle)
			plays++
		case keyboard.KeyEnter:
			if is_taquin_completed(&puzzle) {
				endTime := time.Now().Unix()
				println("Well play ! You solved the puzzle:", ID)
				println("You did", plays, "plays in", endTime-startTime, "seconds")
				taquinSolve = true
			} else {
				println("Puzzle isn't solved yet.")
			}
		case keyboard.KeyEsc:
			taquinSolve = true
		default:
			fmt.Printf("Press ARROW keys to move the puzzle, ENTER to valide your puzzle or ESC to quit.\n")

		}
	}
}

func playCmd(args []string) {
	if len(args) > 0 {
		for _, data := range globalData {
			if data.ID == args[0] {
				cpy := createPuzzleCopy(data.TaquinStruct)
				startGame(data.ID, cpy)
			}
		}
	}
}
