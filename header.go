package main

var globalData []metaTaquin

type metaTaquin struct {
	ID           string
	TaquinStruct taquin
}

type taquin struct {
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

type opti struct{}

type Item struct {
	value    node
	priority uint16
	index    int
}

type PriorityQueue []*Item

type Heuristic func(*taquin) uint16

type Algorithm func(*taquin)

var heur, algo = 0, 0
