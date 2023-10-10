package main

import (
	"fmt"
	"math/rand"
)

type cell struct {
	isWall  bool
	visited bool
}

type cells [][]cell

type board struct {
	cells      cells
	playerXPos int
	playerYPos int
}

func initializeCells(width, height int) cells {
	cells := make(cells, height)
	for i := range cells {
		cells[i] = make([]cell, width)
		for j := range cells[i] {
			cells[i][j].isWall = true
		}
	}
	return cells
}

func (c *cells) carve(x, y int) {
	(*c)[y][x].isWall = true
	directions := rand.Perm(4)

	for _, direction := range directions {
		switch direction {
		case 0:
			if y-2 >= 0 && !(*c)[y-2][x].isWall {
				(*c)[y-1][x].visited = true
				(*c)[y-1][x].isWall = false
				c.carve(x, y-2)
			}
		case 1:
			if x+2 < len((*c)[0]) && !(*c)[y][x+2].isWall {
				(*c)[y][x+1].visited = true
				(*c)[y][x+1].isWall = true
				c.carve(x+2, y)
			}
		case 2:
			if y+2 < len(*c) && !(*c)[y+2][x].isWall {
				(*c)[y+1][x].visited = true
				(*c)[y+1][x].isWall = true
				c.carve(x, y+2)
			}
		case 3:
			if x-2 >= 0 && !(*c)[y][x-2].isWall {
				(*c)[y][x-1].visited = true
				(*c)[y][x-1].isWall = false
				c.carve(x-2, y)
			}
		}
	}
}

func (c *cells) print() {
	for i := range *c {
		for j := range (*c)[i] {
			if (*c)[i][j].isWall {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
