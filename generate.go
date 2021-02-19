package main

import ("fmt"
		"os"
		"strconv"
)

type taquin struct {
	taquin [][]uint16
	size uint8
	voidpos [2]uint8
}

func print_taquin(t taquin) {
	var i uint8

	for i = 0; i < t.size; i++ {
		fmt.Println(t.taquin[i])
	}
}

func generate_taquin(size uint8) taquin {
	var i, j uint8
	var tmp uint16 = 0

	t := taquin{}
	t.taquin = make([][]uint16, size)
	t.size = size
	t.voidpos = [2]uint8{0, 0}
	for i = 0; i < size; i++ {
		t.taquin[i] = make([]uint16, size)
		for j = 0; j < size; j++ {
			t.taquin[i][j] = tmp
			tmp++
		}
	}
	return t
}

func main() {
	i, _ := strconv.ParseInt(os.Args[1], 10, 8);
	t := generate_taquin(uint8(i))
	mix_taquin(&t)
	print_taquin(t)
	fmt.Println("")
	solve(&t)
	print_taquin(t)
}