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

func initializeCells(width, height int) [][]cell {
	cells := make([][]cell, height)
	for i := range cells {
		cells[i] = make([]cell, width)
	}
	return cells
}
