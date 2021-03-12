package main

import ("fmt"
		"os"
		"strconv"
		"time"
)

type taquin struct {
	taquin [][]uint16
	size uint8
	voidpos Vector2D
}

func print_taquin(t taquin) {
	var i uint8

	for i = 0; i < t.size; i++ {
		fmt.Println(t.taquin[i])
	}
}

func generate_taquin(size uint8) taquin {
	var i, j uint8
	var tmp uint16 = 1

	t := taquin{}
	t.taquin = make([][]uint16, size)
	t.size = size
	t.voidpos = Vector2D{size-1, size-1}
	for i = 0; i < size; i++ {
		t.taquin[i] = make([]uint16, size)
		for j = 0; j < size; j++ {
			t.taquin[i][j] = tmp
			tmp++
		}
	}
	t.taquin[size-1][size-1] = 0
	return t
}

func main() {
	i, _ := strconv.ParseInt(os.Args[1], 10, 8);
	t := generate_taquin(uint8(i))
	mix_taquin(&t)
	t2 := copy_taquin(t)
	print_taquin(t)
	fmt.Println("")
	starttime := time.Now()
	taquin_to_string(&t)
	solve(&t)
	solve2(&t2)
	print_taquin(t)
	fmt.Printf("total time: %s\n", time.Since(starttime))
}