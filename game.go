package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	PLAYER = "┃"
	GATE   = "┃"
)

type Difficulty string

const (
	EASY_LEVEL   Difficulty = "easy"
	MEDIUM_LEVEL Difficulty = "medium"
	HARD_LEVEL   Difficulty = "hard"
)

func (b *board) checkWin() bool {
	if b.playerXPos == b.gateXPos && b.playerYPos == b.gateYPos {
		return true
	}
	return false
}

func getDifficultyLevel() string {
	var difficulty string
	validDifficultyLevels := []Difficulty{EASY_LEVEL, MEDIUM_LEVEL, HARD_LEVEL}

	for {
		fmt.Print("Choose a difficulty level (easy, medium, or hard): ")
		fmt.Scanln(&difficulty)
		difficulty = strings.ToLower(difficulty)

		for _, validLevel := range validDifficultyLevels {
			if difficulty == string(validLevel) {
				return difficulty
			}
		}

		fmt.Println("Invalid difficulty level. Please choose from easy, medium, or hard.")
	}
}

func initializeBoard(difficultyLevel string) *board {
	var b *board
	switch difficultyLevel {
	case "hard":
		width, height := 32, 16
		b = &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) * 2) + 1, playerYPos: (rand.Intn(height) * 2) + 1}
	case "medium":
		width, height := 24, 10
		b = &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) * 2) + 1, playerYPos: (rand.Intn(height) * 2) + 1}
	case "easy":
		width, height := 12, 6
		b = &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) * 2) + 1, playerYPos: (rand.Intn(height) * 2) + 1}
	}
	b.cells.carve(b.playerXPos, b.playerYPos)
	gateX, gateY := findFarthestPoint(b.cells, b.playerXPos, b.playerYPos)
	b.gateXPos = gateX
	b.gateYPos = gateY
	b.cells[gateY][gateX].isWall = false
	return b

}
