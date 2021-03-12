package main

import "math/rand"
import "time"

const (
	up = iota
	down = iota
	left = iota
	right = iota
)

func swap(x1 uint8, y1 uint8, x2 uint8, y2 uint8, t *taquin) {
	var tmp uint16
	tmp = t.taquin[y1][x1]
	t.taquin[y1][x1] = t.taquin[y2][x2]
	t.taquin[y2][x2] = tmp
	t.voidpos = Vector2D{x2, y2}
}

func move_left(t *taquin) bool {
	if (t.voidpos.x < t.size-1) {
		swap(t.voidpos.x, t.voidpos.y, t.voidpos.x+1, t.voidpos.y, t)
		return true
	} else {
		return false
	}
}

func move_right(t *taquin) bool {
	if (t.voidpos.x > 0) {
		swap(t.voidpos.x, t.voidpos.y, t.voidpos.x-1, t.voidpos.y, t)
		return true
	} else {
		return false
	}
}

func move_up(t *taquin) bool {
	if (t.voidpos.y < t.size-1) {
		swap(t.voidpos.x, t.voidpos.y, t.voidpos.x, t.voidpos.y+1, t)
		return true
	} else {
		return false
	}
}

func move_down(t *taquin) bool {
	if (t.voidpos.y > 0) {
		swap(t.voidpos.x, t.voidpos.y, t.voidpos.x, t.voidpos.y-1, t)
		return true
	} else {
		return false
	}
}

func do_move(i int, t *taquin) bool {
	if (i == down) {
		return move_down(t)
	} else if (i == up) { 
		return move_up(t)
	} else if (i == right) {
		return move_right(t)
	} else if (i == left) {
		return move_left(t)
	}
	return false
}

func mix_taquin(t *taquin) {
	var move_count uint32 = uint32(t.size)*uint32(t.size)*uint32(t.size)
	var i uint32
	var move, oldmove int

	if t.size < 10 {
		move_count *= 3
	}
	oldmove = 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i = 0; i < move_count; i++ {
		move = rand.Intn(4)
		if oldmove == down && move == up {
			move = down
		} else if (oldmove == up && move == down) {
			move = up
		} else if (oldmove == right && move == left) {
			move = right
		} else if (oldmove == left && move == right) {
			move = left
		}
		for do_move(move, t) == false {
			if move > 0 {
				move--
			} else {
				move = left
			}
		}
		oldmove = move
	}
	for do_move(down, t) {}
	for do_move(right, t) {}
}