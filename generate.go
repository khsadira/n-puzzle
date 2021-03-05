package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
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
	t.Voidpos = Vector2D{0, 0}
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

func start() {
	i, _ := strconv.ParseInt(os.Args[1], 10, 8)
	t := generate_taquin(uint8(i))
	mix_taquin(&t)
	t2 := copy_taquin(t)
	print_taquin(t)
	fmt.Println("")
	starttime := time.Now()
	var wg sync.WaitGroup
	wg.Add(2)
	go solve(&t, &wg)
	go solve2(&t2, &wg)
	wg.Wait()
	print_taquin(t)
	fmt.Printf("total time: %s\n", time.Since(starttime))
}
