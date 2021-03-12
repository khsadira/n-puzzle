package main

import (
	"math/rand"
	"time"
)

const (
	up    = iota
	down  = iota
	left  = iota
	right = iota
)

func swap(x1 uint8, y1 uint8, x2 uint8, y2 uint8, t *taquin) {
	var tmp uint16
	tmp = t.Taquin[y1][x1]
	t.Taquin[y1][x1] = t.Taquin[y2][x2]
	t.Taquin[y2][x2] = tmp
	t.Voidpos = Vector2D{x2, y2}
}

func move_left(t *taquin) bool {
	if t.Voidpos.X < t.Size-1 {
		swap(t.Voidpos.X, t.Voidpos.Y, t.Voidpos.X+1, t.Voidpos.Y, t)
		return true
	} else {
		return false
	}
}

func move_right(t *taquin) bool {
	if t.Voidpos.X > 0 {
		swap(t.Voidpos.X, t.Voidpos.Y, t.Voidpos.X-1, t.Voidpos.Y, t)
		return true
	} else {
		return false
	}
}

func move_up(t *taquin) bool {
	if t.Voidpos.Y < t.Size-1 {
		swap(t.Voidpos.X, t.Voidpos.Y, t.Voidpos.X, t.Voidpos.Y+1, t)
		return true
	} else {
		return false
	}
}

func move_down(t *taquin) bool {
	if t.Voidpos.Y > 0 {
		swap(t.Voidpos.X, t.Voidpos.Y, t.Voidpos.X, t.Voidpos.Y-1, t)
		return true
	} else {
		return false
	}
}

func do_move(i int, t *taquin) bool {
	if i == down {
		return move_down(t)
	} else if i == up {
		return move_up(t)
	} else if i == right {
		return move_right(t)
	} else if i == left {
		return move_left(t)
	}
	return false
}

func mix_taquin(t *taquin) {
	var move_count uint32 = uint32(t.Size) * uint32(t.Size) * uint32(t.Size)
	var i uint32
	var move, oldmove int

	if t.Size < 10 {
		move_count *= 3
	}
	oldmove = 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i = 0; i < move_count; i++ {
		move = rand.Intn(4)
		if oldmove == down && move == up {
			move = down
		} else if oldmove == up && move == down {
			move = up
		} else if oldmove == right && move == left {
			move = right
		} else if oldmove == left && move == right {
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
	for do_move(down, t) {
	}
	for do_move(right, t) {
	}
}
