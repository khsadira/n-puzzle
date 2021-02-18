package main

import "math/rand"
import "time"

func swap(x1 uint8, y1 uint8, x2 uint8, y2 uint8, t *taquin) {
	var tmp uint16
	tmp = t.taquin[y1][x1]
	t.taquin[y1][x1] = t.taquin[y2][x2]
	t.taquin[y2][x2] = tmp
}

func move_left(t *taquin) bool {
	if (t.voidpos[1] < t.size-1) {
		swap(t.voidpos[1], t.voidpos[0], t.voidpos[1]+1, t.voidpos[0], t)
		t.voidpos[1]++
		return true
	} else {
		return false
	}
}

func move_right(t *taquin) bool {
	if (t.voidpos[1] > 0) {
		swap(t.voidpos[1], t.voidpos[0], t.voidpos[1]-1, t.voidpos[0], t)
		t.voidpos[1]--
		return true
	} else {
		return false
	}
}

func move_up(t *taquin) bool {
	if (t.voidpos[0] < t.size-1) {
		swap(t.voidpos[1], t.voidpos[0], t.voidpos[1], t.voidpos[0]+1, t)
		t.voidpos[0]++
		return true
	} else {
		return false
	}
}

func move_down(t *taquin) bool {
	if (t.voidpos[0] > 0) {
		swap(t.voidpos[1], t.voidpos[0], t.voidpos[1], t.voidpos[0]-1, t)
		t.voidpos[0]--
		return true
	} else {
		return false
	}
}

func do_move(i uint8, t *taquin) bool {
	if (i == 0) {
		return move_down(t)
	} else if (i == 1) { 
		return move_up(t)
	} else if (i == 2) {
		return move_right(t)
	} else if (i == 3) {
		return move_left(t)
	}
	return false
}

func mix_taquin(t *taquin) {
	var move_count uint32 = uint32(t.size)*uint32(t.size)*uint32(t.size)
	var i uint32
	var move, oldmove uint8

	if t.size < 10 {
		move_count *= 3
	}
	oldmove = 0
	rand.Seed(time.Now().UTC().UnixNano())
	for i = 0; i < move_count; i++ {
		move = uint8(rand.Intn(4))
		if oldmove == 0 && move == 1 {
			move = 0
		} else if (oldmove == 1 && move == 0) {
			move = 1
		} else if (oldmove == 2 && move == 3) {
			move = 2
		} else if (oldmove == 3 && move == 2) {
			move = 3
		}
		for do_move(move, t) == false {
			if move > 0 {
				move--
			} else {
				move = 3
			}
		}
		oldmove = move
	}
	for move_down(t) {}
	for move_right(t) {}
}