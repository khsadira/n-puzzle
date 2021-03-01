package main

type taquin struct {
	ID      string
	taquin  [][]uint16
	size    uint8
	voidpos Vector2D
}

type Vector2D struct {
	x, y uint8
}

type node struct {
	pos             Vector2D
	cost, heuristic uint16
	parent_move     int
	t               taquin
	parent_node     *node
}
