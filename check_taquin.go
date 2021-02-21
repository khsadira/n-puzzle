package main

func convertTaquinToArray(puzzle taquin) []uint16 {
	var taquinArray []uint16

	for _, line := range puzzle.taquin {
		for _, value := range line {
			taquinArray = append(taquinArray, value)
		}
	}

	return taquinArray
}

func checkTaquin(puzzle taquin) bool {
	var taquinArray []uint16 = convertTaquinToArray(puzzle)
	var inversion uint16 = 0

	for i := range taquinArray {
		var currentInversion uint16 = taquinArray[i]
		for j := 0; j <= i; j++ {
			if taquinArray[j] != 0 && taquinArray[i] >= taquinArray[j] && currentInversion > 0 {
				currentInversion--
			}
		}
		inversion += currentInversion
	}
	if inversion%2 == 1 {
		println(puzzle.ID, "is unsolvable.")
		return false
	}
	return true
}
