package main

import (
	"fmt"
)

func print_taquin(t taquin) {
	var i uint8

	for i = 0; i < t.Size; i++ {
		fmt.Println(t.Taquin[i])
	}
}

func generate_taquin(size uint8) taquin {
	var i, j uint8
	var tmp uint16 = 1

	t := taquin{}
	t.Taquin = make([][]uint16, size)
	t.Size = size
	t.Voidpos = Vector2D{size - 1, size - 1}
	for i = 0; i < size; i++ {
		t.Taquin[i] = make([]uint16, size)
		for j = 0; j < size; j++ {
			t.Taquin[i][j] = tmp
			tmp++
		}
	}
	t.Taquin[size-1][size-1] = 0
	return t
}
