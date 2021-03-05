package main

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"sync"
	"time"
)

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

func get_target_pos(val uint16, t *taquin) Vector2D {
	if val == 0 {
		return Vector2D{t.Size - 1, t.Size - 1}
	}
	var x, y uint8
	x = uint8(val % uint16(t.Size))
	if x == 0 {
		x = t.Size
	}
	y = uint8(math.Ceil(float64(val) / float64(t.Size)))
	return Vector2D{x - 1, y - 1}
}

func calc_manhattan_distance(p1 Vector2D, p2 Vector2D) uint16 {
	return uint16(math.Abs(float64(p1.X)-float64(p2.X)) + math.Abs(float64(p1.Y)-float64(p2.Y)))
}

func calc_heuristic_manhattan_distance(t *taquin) uint16 {
	var ret uint16 = 0
	var i, j uint8
	var val uint16 = 1

	for i = 0; i < t.Size; i++ {
		for j = 0; j < t.Size; j++ {
			if t.Taquin[i][j] != val && t.Taquin[i][j] != 0 {
				ret += calc_manhattan_distance(Vector2D{j, i}, get_target_pos(t.Taquin[i][j], t))
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

	for i = 0; i < t.Size; i++ {
		for j = 0; j < t.Size; j++ {
			if t.Taquin[i][j] != val && t.Taquin[i][j] != 0 {
				ret++
			}
			val++
		}
	}
	return ret
}

func not_reverse_move(move int, parent_move int) bool {
	if (move == up && parent_move == down) ||
		(move == down && parent_move == up) ||
		(move == right && parent_move == left) ||
		(move == left && parent_move == right) {
		return false
	}
	return true
}

func get_reverse_move(move int) int {
	if move == right {
		return left
	} else if move == left {
		return right
	} else if move == up {
		return down
	} else if move == down {
		return up
	}
	return -1
}

func remove_node_from_slice(list []node, n node) []node {
	var i int

	for i = 0; i < len(list); i++ {
		if are_taquins_equal(&n.t, &list[i].t) {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func are_taquins_equal(t1 *taquin, t2 *taquin) bool {
	var i, j uint8

	for i = 0; i < t1.Size; i++ {
		for j = 0; j < t1.Size; j++ {
			if t1.Taquin[i][j] != t2.Taquin[i][j] {
				return false
			}
		}
	}
	return true
}

func is_node_in_slice(list []node, n node) bool {
	var i int

	for i = 0; i < len(list); i++ {
		if are_taquins_equal(&n.t, &list[i].t) {
			return true
		}
	}
	return false
}

func is_node_in_slice_with_less_cost(list []node, n node) bool {
	var i int

	for i = 0; i < len(list); i++ {
		if are_taquins_equal(&n.t, &list[i].t) && n.cost < list[i].cost {
			return true
		}
	}
	return false
}

func copy_taquin(t taquin) taquin {
	var ret taquin
	var i uint8

	ret.Voidpos = Vector2D{t.Voidpos.X, t.Voidpos.Y}
	ret.Size = t.Size
	ret.Taquin = make([][]uint16, len(t.Taquin))
	for i = 0; i < ret.Size; i++ {
		ret.Taquin[i] = make([]uint16, len(t.Taquin[i]))
		copy(ret.Taquin[i], t.Taquin[i])
	}
	return ret
}

func solve(t *taquin, wg *sync.WaitGroup) {
	var open_list []node
	var close_list []node
	var n, newn node
	var i int

	start := time.Now()
	n = node{t.Voidpos, 0, calc_heuristic_manhattan_distance(t), -1, copy_taquin(*t), nil}
	open_list = append(open_list, n)
	for n.heuristic != 0 {
		for i = 0; i < 4; i++ {
			if not_reverse_move(i, n.parent_move) && do_move(i, t) {
				newn = node{t.Voidpos, n.cost + 1, calc_heuristic_manhattan_distance(t), i, copy_taquin(*t), &n}
				if newn.heuristic == 0 {
					fmt.Printf("%s: %d\n", "cost", n.cost)
					PrintMemUsage()
					fmt.Println(time.Since(start))
					wg.Done()
					return
				}
				if !(is_node_in_slice(close_list, newn) || is_node_in_slice_with_less_cost(open_list, newn)) {
					open_list = append(open_list, newn)
				}
				do_move(get_reverse_move(i), t)
			}
		}
		close_list = append(close_list, n)
		open_list = open_list[1:]
		sort.Slice(open_list, func(i, j int) bool {
			return (open_list[i].heuristic + open_list[i].cost) < (open_list[j].heuristic + open_list[j].cost)
		})
		n = open_list[0]
		*t = copy_taquin(n.t)
	}
	fmt.Printf("%s: %d\n", "cost", n.cost)
	PrintMemUsage()
	fmt.Println(time.Since(start))
	wg.Done()
}

func solve2(t *taquin, wg *sync.WaitGroup) {
	var open_list []node
	var close_list []node
	var n, newn node
	var i int

	start := time.Now()
	n = node{t.Voidpos, 0, calc_heuristic_manhattan_distance(t), -1, copy_taquin(*t), nil}
	open_list = append(open_list, n)
	for n.heuristic != 0 {
		for i = 0; i < 4; i++ {
			if not_reverse_move(i, n.parent_move) && do_move(i, t) {
				newn = node{t.Voidpos, n.cost + 1, calc_heuristic_manhattan_distance(t), i, copy_taquin(*t), &n}
				if newn.heuristic == 0 {
					fmt.Printf("%s: %d\n", "cost", n.cost)
					PrintMemUsage()
					fmt.Println(time.Since(start))
					wg.Done()
					return
				}
				if !(is_node_in_slice(close_list, newn) || is_node_in_slice_with_less_cost(open_list, newn)) {
					open_list = append(open_list, newn)
				}
				do_move(get_reverse_move(i), t)
			}
		}
		close_list = append(close_list, n)
		open_list = open_list[1:]
		sort.Slice(open_list, func(i, j int) bool {
			return (open_list[i].heuristic + open_list[i].cost) < (open_list[j].heuristic + open_list[j].cost)
		})
		n = open_list[0]
		*t = copy_taquin(n.t)
	}
	fmt.Printf("%s: %d\n", "cost", n.cost)
	PrintMemUsage()
	fmt.Println(time.Since(start))
	wg.Done()
}
