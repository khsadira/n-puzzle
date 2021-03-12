package main

func convertTaquinToArray(taquin [][]uint16) []uint16 {
	var taquinArray []uint16

	for _, line := range taquin {
		for _, value := range line {
			taquinArray = append(taquinArray, value)
		}
	}

	return taquinArray
}

func checkTaquinValues(taquinArray []uint16) bool {
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

	if voidPosRaw%2 == 1 && inversion%2 == 0 || voidPosRaw%2 == 0 && inversion%2 == 1 {
		return true
	}

	return false
}

func isValidTaquin(ID string, puzzle taquin) bool {
	var taquinArray []uint16 = convertTaquinToArray(puzzle.Taquin)
	var inversion uint16 = getInversionNumber(taquinArray)

	if !checkTaquinValues(taquinArray) {
		println("n-puzzle:", ID, "values are incorrect.")
		return false
	}

	if puzzle.Size%2 == 1 && !checkOddTaquin(inversion) || puzzle.Size%2 == 0 && !checkEvenTaquin(puzzle, inversion) {
		println("n-puzzle:", ID, "is unsolvable.")
		return false
	}

	return true
}
