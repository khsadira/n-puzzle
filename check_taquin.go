package main

import "fmt"

func convertTaquinToArray(taquin [][]uint16) []uint16 {
	var taquinArray []uint16

	for _, line := range taquin {
		for _, value := range line {
			taquinArray = append(taquinArray, value)
		}
	}

	return taquinArray
}

func isValidTaquin(puzzle taquin) bool {
	var taquinArray []uint16 = convertTaquinToArray(puzzle.taquin)

	for i := range taquinArray {
		for j := 0; j < i; j++ {
			if taquinArray[i] == taquinArray[j] {
				println("n-puzzle:", puzzle.ID, "is unsolvable.1")
				return false
			}
		}
	}

	//https://www.geeksforgeeks.org/check-instance-15-puzzle-solvable/#:~:text=In%20general%2C%20for%20a%20given,even%20in%20the%20input%20state.&text=the%20blank%20is%20on%20an,fourth%2Dlast%2C%20etc.)
	var inversion uint16 = 0
	taquinLen := len(taquinArray)

	for i := 0; i < taquinLen-1; i++ {
		for j := i + 1; j < taquinLen; j++ {
			if taquinArray[i] != 0 && taquinArray[j] != 0 && taquinArray[i] > taquinArray[j] {
				inversion++
			}
		}
	}
	println(inversion)
	if inversion%2 == 1 {
		println("n-puzzle:", puzzle.ID, "is unsolvable.2")
		return false
	}

	for i := range taquinArray {
		for j := 0; j < i; j++ {
			if taquinArray[i] < taquinArray[j] {
				taquinArray[i], taquinArray[j] = taquinArray[j], taquinArray[i]
			}
		}
	}

	for key, value := range taquinArray {
		if uint16(key) != value {
			println("n-puzzle:", puzzle.ID, "is unsolvable or already solved.3")
			return false
		}
	}

	fmt.Printf("%v\n", taquinArray)
	return true
}
