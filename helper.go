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
				voidpos.Y = i
				voidpos.X = j
				return voidpos, nil
			}
		}
	}
	return voidpos, errors.New("puzzle not well formatted.")
}

func appendDataToGlobalData(dataToAdd metaTaquin) {
	for _, data := range globalData {
		if data.ID == dataToAdd.ID {
			println("n-puzzle: error: Failed to add", dataToAdd.ID, "due to a similar ID finded.")
			return
		}
	}

	globalData = append(globalData, dataToAdd)
	println("n-puzzle:", dataToAdd.ID, "added to global data puzzles.")
}

func showPuzzle(puzzle taquin) {
	for i := uint8(0); i < puzzle.Size; i++ {
		var IDs []string
		for j := uint8(0); j < puzzle.Size; j++ {
			IDs = append(IDs, strconv.Itoa(int(puzzle.Taquin[i][j])))
		}
		fmt.Println(strings.Join(IDs, " "))
	}
}

func isTaquinSolved(puzzle taquin) bool {
	var taquinArray []uint16 = convertTaquinToArray(puzzle.Taquin)
	jmp := 0

	fmt.Printf("%v\n", taquinArray)
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

func removeData(ID string) {
	for i, data := range globalData {
		if data.ID == ID {
			globalData = append(globalData[:i], globalData[i+1:]...)
		}
	}
}

func createPuzzleCopy(t taquin) taquin {
	var cpy taquin

	cpy.Size = t.Size
	cpy.Voidpos = Vector2D{
		X: t.Voidpos.X,
		Y: t.Voidpos.Y}

	cpy.Taquin = make([][]uint16, len(t.Taquin))
	for i := 0; i < len(t.Taquin); i++ {
		cpy.Taquin[i] = make([]uint16, len(t.Taquin[i]))
		for j := 0; j < len(t.Taquin[i]); j++ {
			cpy.Taquin[i][j] = t.Taquin[i][j]
		}
	}

	return cpy
}
