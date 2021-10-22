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
	var offset uint8 = 0

	t := taquin{}
	t.Taquin = make([][]uint16, size)
	t.Size = size
	t.Voidpos = Vector2D{uint8((size-1) / 2), uint8(size / 2)}
	for i = 0; i < size; i++ {
		t.Taquin[i] = make([]uint16, size)
	}
	i = 0
	j = 0
	for i != t.Voidpos.Y || j != t.Voidpos.X {
		if i == offset && j == offset {
			for j < size-offset-1 {
				t.Taquin[i][j] = tmp
				tmp++
				j++
			}
		} else if i == offset && j == size-1-offset {
			for i < size-offset-1 {
				t.Taquin[i][j] = tmp
				tmp++
				i++
			}
		} else if i == size-1-offset && j == size-1-offset {
			for j > offset {
				t.Taquin[i][j] = tmp
				tmp++
				j--
			}
		} else if i == size-1-offset && j == offset {
			for i > offset {
				t.Taquin[i][j] = tmp
				tmp++
				i--
			}
			offset++
			i = offset
			j = offset
		}
	}
	t.Taquin[t.Voidpos.Y][t.Voidpos.X] = 0
	return t
}
