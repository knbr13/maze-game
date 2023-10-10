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

const (
	PLAYER = "┃"
)

func initializeCells(width, height int) cells {
	cells := make(cells, height*2+1)
	for i := range cells {
		cells[i] = make([]cell, width*2+1)
		for j := range cells[i] {
			cells[i][j].isWall = true
		}
	}
	return cells
}

func (c *cells) carve(x, y int) {
	(*c)[y][x].visited = true
	(*c)[y][x].isWall = false
	directions := rand.Perm(4)

	for _, direction := range directions {
		switch direction {
		case 0:
			if y-2 > 0 && !(*c)[y-2][x].visited {
				(*c)[y-1][x].isWall = false
				c.carve(x, y-2)
			}
		case 1:
			if x+2 < len((*c)[y])-1 && !(*c)[y][x+2].visited {
				(*c)[y][x+1].isWall = false
				c.carve(x+2, y)
			}
		case 2:
			if y+2 < len(*c)-1 && !(*c)[y+2][x].visited {
				(*c)[y+1][x].isWall = false
				c.carve(x, y+2)
			}
		case 3:
			if x-2 > 0 && !(*c)[y][x-2].visited {
				(*c)[y][x-1].isWall = false
				c.carve(x-2, y)
			}
		}
	}
}

func (b *board) print() {
	for i, row := range b.cells {
		for j, col := range row {
			if i == b.playerYPos && j == b.playerXPos {
				fmt.Print(PLAYER)
			} else if col.isWall {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (b *board) handleMove(r rune) {
	switch r {
	case 'w':
		if b.playerYPos > 0 && !b.cells[b.playerYPos-1][b.playerXPos].isWall {
			b.playerYPos--
		}
	case 'd':
		if b.playerXPos < len(b.cells[b.playerYPos])-1 && !b.cells[b.playerYPos][b.playerXPos+1].isWall {
			b.playerXPos++
		}
	case 's':
		if b.playerYPos < len(b.cells)-1 && !b.cells[b.playerYPos+1][b.playerXPos].isWall {
			b.playerYPos++
		}
	case 'a':
		if b.playerXPos > 0 && !b.cells[b.playerYPos][b.playerXPos-1].isWall {
			b.playerXPos--
		}
	}
}
