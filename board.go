package main

type cell struct {
	isWall  bool
	visited bool
}

type board struct {
	cells      [][]cell
	playerXPos int
	playerYPos int
}
