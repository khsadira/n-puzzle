package main

import "fmt"
import "runtime"
import "math"
import "time"
import "container/heap"
import "strings"

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

type node struct {
	pos Vector2D
	cost, heuristic uint16
	parent_move int
	t taquin
	parent_node *node
}

type opti struct {}

type Vector2D struct {
	x, y uint8
}

func get_target_pos(val uint16, t *taquin) Vector2D {
	if (val == 0) {
		return Vector2D{t.size-1, t.size-1}
	}
	var x, y uint8
	x = uint8(val % uint16(t.size))
	if (x == 0) {
		x = t.size
	}
	y = uint8(math.Ceil(float64(val) / float64(t.size)))
	return Vector2D{x-1, y-1}
}

func calc_manhattan_distance(p1 Vector2D, p2 Vector2D) uint16 {
	return uint16(math.Abs(float64(p1.x) - float64(p2.x)) + math.Abs(float64(p1.y) - float64(p2.y)))
}

func calc_heuristic_manhattan_distance(t *taquin) uint16 {
	var ret uint16 = 0
	var i, j uint8
	var val uint16 = 1

	for i = 0; i < t.size; i++ {
		for j = 0; j < t.size; j++ {
			if (t.taquin[i][j] != val && t.taquin[i][j] != 0) {
				ret += calc_manhattan_distance(Vector2D{j, i}, get_target_pos(t.taquin[i][j], t))
			}
			val++
		}
	}
	return ret
}

func calc_heuristic_nb_misplaced(t *taquin) uint16 {
	var ret uint16 = 0
	var i, j uint8
	var val uint16 = 1

	for i = 0; i < t.size; i++ {
		for j = 0; j < t.size; j++ {
			if (t.taquin[i][j] != val && t.taquin[i][j] != 0) {
				ret++
			}
			val++
		}
	}
	return ret
}

func not_reverse_move(move int, parent_move int) bool {
	if ((move == up && parent_move == down) ||
			(move == down && parent_move == up) ||
			(move == right && parent_move == left) ||
			(move == left && parent_move == right)) {
		return false
	}
	return true
}

func get_reverse_move(move int) int {
	if (move == right) {
		return left
	} else if (move == left) {
		return right
	} else if (move == up) {
		return down
	} else if (move == down) {
		return up
	}
	return -1
}

func copy_taquin(t taquin) taquin {
	var ret taquin
	var i uint8

	ret.voidpos = Vector2D{t.voidpos.x, t.voidpos.y}
	ret.size = t.size
	ret.taquin = make([][]uint16, len(t.taquin))
	for i = 0; i < ret.size; i++ {
		ret.taquin[i] = make([]uint16, len(t.taquin[i]))
		copy(ret.taquin[i], t.taquin[i])
	}
	return ret
}

func taquin_to_string(t *taquin) string {
	var i int
	var ret string = ""

	for i = 0; i < int(t.size); i++ {
		ret += strings.Trim(strings.Replace(fmt.Sprint(t.taquin[i]), " ", ",", -1), "[]")
		if (i < int(t.size) - 1) {
			ret += ","
		}
	}
	return ret
}

func solve(t *taquin) {
	var n, newn node
	var i int
	var newItem *Item

	close_list := make(map[string]opti)
	open_list := make(PriorityQueue, 1)
	start := time.Now()
	n = node{t.voidpos, 0, calc_heuristic_manhattan_distance(t), -1, copy_taquin(*t), nil}
	open_list[0] = &Item{n, n.heuristic + n.cost, 0}
	heap.Init(&open_list)
	for n.heuristic != 0 {
		n = heap.Pop(&open_list).(*Item).value
		*t = copy_taquin(n.t)
		for i = 0; i < 4; i++ {
			if (not_reverse_move(i, n.parent_move) && do_move(i, t)) {
				newn = node{t.voidpos, n.cost+1, calc_heuristic_manhattan_distance(t), i, copy_taquin(*t), &n}
				if (newn.heuristic == 0) {
					fmt.Printf("%s: %d\n", "cost", n.cost)
					PrintMemUsage()
					fmt.Println(time.Since(start))
					return
				}
				_, ok := close_list[taquin_to_string(&newn.t)]
				if !(ok) {
					newItem = &Item{newn, newn.heuristic + n.cost, 0}
					heap.Push(&open_list, newItem)
				}
				do_move(get_reverse_move(i), t)
			}
		}
		close_list[taquin_to_string(&n.t)] = opti{}
	}
}

func solve2(t *taquin) {
	var n, newn node
	var i int
	var newItem *Item

	close_list := make(map[string]opti)
	open_list := make(PriorityQueue, 1)
	start := time.Now()
	n = node{t.voidpos, 0, calc_heuristic_manhattan_distance(t), -1, copy_taquin(*t), nil}
	open_list[0] = &Item{n, n.heuristic + n.cost, 0}
	heap.Init(&open_list)
	for n.heuristic != 0 {
		n = heap.Pop(&open_list).(*Item).value
		*t = copy_taquin(n.t)
		for i = 0; i < 4; i++ {
			if (not_reverse_move(i, n.parent_move) && do_move(i, t)) {
				newn = node{t.voidpos, n.cost+1, calc_heuristic_manhattan_distance(t), i, copy_taquin(*t), &n}
				if (newn.heuristic == 0) {
					fmt.Printf("%s: %d\n", "cost", n.cost)
					PrintMemUsage()
					fmt.Println(time.Since(start))
					return
				}
				_, ok := close_list[taquin_to_string(&newn.t)]
				if !(ok) {
					newItem = &Item{newn, newn.heuristic + n.cost, 0}
					heap.Push(&open_list, newItem)
				}
				do_move(get_reverse_move(i), t)
			}
		}
		close_list[taquin_to_string(&n.t)] = opti{}
	}
}