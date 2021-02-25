package main

import "sort"
import "fmt"
import "runtime"
import "math"

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}

func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
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
	return uint16(math.Abs(float64(p1.x - p2.x)) + math.Abs(float64(p1.y - p2.y)))
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

	for i = 0; i < t1.size; i++ {
		for j = 0; j < t1.size; j++ {
			if (t1.taquin[i][j] != t2.taquin[i][j]) {
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

func algo(t *taquin, open_list []node, close_list []node, current_node node) bool {
	var i int
	var new_node node

	if current_node.heuristic == 0.0 {
		return true
	}
	close_list = append(close_list, current_node)
	open_list = remove_node_from_slice(open_list, current_node)
	for i = 0; i < 4; i++ {
		if not_reverse_move(i, current_node.parent_move) {
			if (do_move(i, t)) {
				new_node = node{t.voidpos, current_node.cost + 1, calc_heuristic_manhattan_distance(t), i, copy_taquin(*t), &current_node}
				do_move(get_reverse_move(i), t)
				if (!is_node_in_slice(close_list, new_node)) {
					open_list = append(open_list, new_node)
				}
			}
		}
	}
	sort.Slice(open_list, func(i, j int) bool {
		return (open_list[i].heuristic+open_list[i].cost) < (open_list[j].heuristic+open_list[j].cost)
	})
	for i = 0; i < len(open_list); i++ {
		*t = copy_taquin(open_list[i].t)
		fmt.Println(len(open_list))
		if (algo(t, open_list, close_list, open_list[i])) {
			return true
		}
	}
	return false
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

func solve(t *taquin) {
	var open_list []node
	var close_list []node
	var n node

	n.cost = 0
	n.heuristic = calc_heuristic_manhattan_distance(t)
	n.parent_move = -1
	n.pos = t.voidpos
	n.t = copy_taquin(*t)
	open_list = append(open_list, n)
	PrintMemUsage()
	algo(t, open_list, close_list, n)
	PrintMemUsage()
}