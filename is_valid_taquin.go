package main

func FindIndexSlice(slice []uint16, value uint16) int {
	for p, v := range slice {
		if value == v {
			return p
		}
	}
	return -1
}
func countInversions(puzzle []uint16) int {
	inversions := 0
	for i := 0; i < len(puzzle)-1; i++ {
		for j := i + 1; j < len(puzzle); j++ {
			if puzzle[i] > puzzle[j] && puzzle[i] != 0 && puzzle[j] != 0 {
				inversions++
			}
		}
	}
	return inversions
}

func oddSize(solution []uint16, puzzle []uint16, size int) bool {
	pInversions := countInversions(puzzle)
	sInversions := countInversions(solution)
	pIdx := FindIndexSlice(puzzle, 0)
	sIdx := FindIndexSlice(solution, 0)

	if size%2 == 0 {
		pInversions += pIdx / size
		sInversions += sIdx / size
	}

	return ((pInversions % 2) == (sInversions % 2))
}

func evenSize(solution []uint16, puzzle []uint16, size int) bool {
	Inversions := countInversions(puzzle)
	zeroIdx := FindIndexSlice(puzzle, 0)
	row := (((size*size - 1) - zeroIdx) / size) + 1
	if ((row%2 == 0) && (Inversions%2 != 0)) || (row%2 != 0) && (Inversions%2 == 0) {
		return true
	}
	return false
}

func IsSolvable(solution []uint16, puzzle []uint16, size int) bool {
	if size%2 != 0 || size == 6 || size == 8 {
		return oddSize(solution, puzzle, size)
	} else {
		return evenSize(solution, puzzle, size)
	}
}

func convertTaquinToArray(taquin [][]uint16) []uint16 {
	var taquinArray []uint16

	for _, line := range taquin {
		for _, value := range line {
			taquinArray = append(taquinArray, value)
		}
	}
	return taquinArray
}

func checkTaquinValues(start []uint16) bool {
	taquinArray := make([]uint16, len(start))
	copy(taquinArray, start)

	for i := range taquinArray {
		for j := 0; j < i; j++ {
			if taquinArray[i] < taquinArray[j] {
				taquinArray[i], taquinArray[j] = taquinArray[j], taquinArray[i]
			}
		}
	}

	for i := 0; i < len(taquinArray); i++ {
		if taquinArray[i] != uint16(i) {
			return false
		}
	}

	return true
}

func getInversionNumber(taquinArray []uint16) uint16 {
	var inversion uint16 = 0
	taquinLen := len(taquinArray)

	for i := 0; i < taquinLen-1; i++ {
		for j := i + 1; j < taquinLen; j++ {
			if taquinArray[i] != 0 && taquinArray[j] != 0 && taquinArray[i] > taquinArray[j] {
				inversion++
			}
		}
	}

	return inversion
}

func checkOddTaquin(inversion uint16) bool {
	if inversion%2 == 1 {
		return false
	}
	return true
}

func checkEvenTaquin(puzzle taquin, inversion uint16) bool {
	var voidPosRaw uint8 = puzzle.Voidpos.Y

	println(voidPosRaw)

	if voidPosRaw%2 == 1 && inversion%2 == 0 || voidPosRaw%2 == 0 && inversion%2 == 1 {
		return true
	}

	return false
}

func isValidTaquin(ID string, puzzle taquin) bool {
	var solutionPuzzle = generate_taquin(puzzle.Size)
	taquinArray := convertTaquinToArray(puzzle.Taquin)
	solutionArray := convertTaquinToArray(solutionPuzzle.Taquin)

	if !checkTaquinValues(taquinArray) {
		println("n-puzzle:", ID, "values are incorrect.")
		return false
	}

	if !IsSolvable(solutionArray, taquinArray, int(puzzle.Size)) {
		println("n-puzzle:", ID, "is unsolvable.")
		return false
	}

	return true
}
