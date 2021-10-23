package main

func get_snail_taquin(puzzle [][]uint16, size uint8) []uint16 {

	r := size
	c := size

	left := 0
	right := int(c - 1)

	top := 0
	bottom := int(r - 1)
	dir := 0
	var ret []uint16

	for left <= right && top <= bottom {
		if dir == 0 {
			for i := left; i <= right; i++ {
				ret = append(ret, puzzle[top][i])
			}
			top++
			dir = 1
		} else if dir == 1 {
			for i := top; i <= bottom; i++ {
				ret = append(ret, puzzle[i][right])
			}
			right--
			dir = 2
		} else if dir == 2 {
			for i := right; i >= left; i-- {
				ret = append(ret, puzzle[bottom][i])
			}
			bottom--
			dir = 3
		} else if dir == 3 {
			for i := bottom; i >= top; i-- {
				ret = append(ret, puzzle[i][left])
			}
			left++
			dir = 0
		}
	}

	return ret
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
	var taquinArray = get_snail_taquin(puzzle.Taquin, puzzle.Size)
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
