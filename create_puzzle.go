package main

func createPuzzleTest(ID string, size uint8) taquin {
	var puzzle taquin

	nPuzzle := make([][]uint16, size)
	puzzle.Size = size
	puzzle.ID = ID

	puzzle.Voidpos.X = 0
	puzzle.Voidpos.Y = 0
	for i := uint8(0); i < size; i++ {
		nPuzzle[i] = make([]uint16, size)
		for j := uint8(0); j < size; j++ {
			nPuzzle[i][j] = uint16(i*size + j)
		}
	}

	puzzle.Taquin = nPuzzle

	return puzzle
}
