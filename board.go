package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/eiannone/keyboard"
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
	gateXPos   int
	gateYPos   int
}

const (
	PLAYER = "┃"
	GATE   = "G"
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

func (b *board) String() string {
	s := ""
	for i, row := range b.cells {
		for j, col := range row {
			if i == b.playerYPos && j == b.playerXPos {
				s += fmt.Sprint(PLAYER)
			} else if col.isWall {
				s += "#"
			} else if i == b.gateYPos && j == b.gateXPos {
				s += fmt.Sprint(GATE)
			} else {
				s += " "
			}
		}
		s += fmt.Sprintln()
	}
	return s
}

func (b *board) handleMove(r rune, key keyboard.Key) {
	switch {
	case r == 'q':
		os.Exit(0)
	case r == 'w' || key == keyboard.KeyArrowUp:
		if b.playerYPos > 0 && !b.cells[b.playerYPos-1][b.playerXPos].isWall {
			b.playerYPos--
		}
	case r == 'd' || key == keyboard.KeyArrowRight:
		if b.playerXPos < len(b.cells[b.playerYPos])-1 && !b.cells[b.playerYPos][b.playerXPos+1].isWall {
			b.playerXPos++
		}
	case r == 's' || key == keyboard.KeyArrowDown:
		if b.playerYPos < len(b.cells)-1 && !b.cells[b.playerYPos+1][b.playerXPos].isWall {
			b.playerYPos++
		}
	case r == 'a' || key == keyboard.KeyArrowLeft:
		if b.playerXPos > 0 && !b.cells[b.playerYPos][b.playerXPos-1].isWall {
			b.playerXPos--
		}
	}
}

func findFarthestPoint(c cells, startX, startY int) (int, int) {
	distances := make([][]int, len(c))
	for i := range distances {
		distances[i] = make([]int, len(c[i]))
		for j := range distances[i] {
			distances[i][j] = -1
		}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{startX, startY})
	distances[startY][startX] = 0

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]

		for _, dir := range []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX, newY := x+dir.dx, y+dir.dy
			if newX >= 0 && newX < len(c[0]) && newY >= 0 && newY < len(c) && !c[newY][newX].isWall && distances[newY][newX] == -1 {
				distances[newY][newX] = distances[y][x] + 1
				queue = append(queue, [2]int{newX, newY})
			}
		}
	}

	maxDistance := -1
	var farthestX, farthestY int

	for y, row := range distances {
		for x, dist := range row {
			if dist > maxDistance {
				maxDistance = dist
				farthestX, farthestY = x, y
			}
		}
	}

	return farthestX, farthestY
}
