package main

import "math"
import "sort"

type node struct {
	pos Vector2D
	cost, heuristic float64
	parent_move int
	t taquin
	parent_node *node
}

type Vector2D struct {
	x, y uint8
}

func distance_to_recquired_pos(cur Vector2D, tar Vector2D) float64 { // perfs a tester, simplifiable
	var x1, y1, x2, y2 float64 = float64(cur.x), float64(cur.y), float64(tar.x), float64(tar.y)
	return math.Sqrt(((x1-x2)*(x1-x2))+((y1-y2)*(y1-y2)))
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

func calc_heuristic(t *taquin) float64 { // a opti
	var x, y uint8
	var h, c float64 = 0, 0

	for y = 0; y < t.size; y++ {
		for x = 0; x < t.size; x++ {
			h += distance_to_recquired_pos(Vector2D{x, y}, get_target_pos(t.taquin[y][x], t))
			c++
		}
	}
	return h
}

func calc_base_heuristic(t *taquin) float64 {
	var x, y uint8
	var h, c float64 = 0, 0

	for y = 0; y < t.size; y++ {
		for x = 0; x < t.size; x++ {
			h += distance_to_recquired_pos(Vector2D{x, y}, get_target_pos(t.taquin[y][x], t))
			c++
		}
	}
	return h
}

func calc_cost(t *taquin, move int) float64 {
	return 2.0
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
				new_node = node{t.voidpos, calc_cost(t, i), calc_heuristic(t), i, copy_taquin(*t), &current_node}
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
		do_move(open_list[i].parent_move, t)
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
	n.heuristic = calc_base_heuristic(t)
	n.parent_move = -1
	n.pos = t.voidpos
	n.t = copy_taquin(*t)
	open_list = append(open_list, n)
	algo(t, open_list, close_list, n)
}