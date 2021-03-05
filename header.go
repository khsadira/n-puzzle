package main

type taquin struct {
	ID      string
	Taquin  [][]uint16
	Size    uint8
	Voidpos Vector2D
}

type Vector2D struct {
	X, Y uint8
}

type node struct {
	pos             Vector2D
	cost, heuristic uint16
	parent_move     int
	t               taquin
	parent_node     *node
}
