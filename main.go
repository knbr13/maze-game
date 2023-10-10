package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	difficultyLevel := getDifficultyLevel()
	board := initializeBoard(difficultyLevel)
	board.cells.carve(board.playerXPos, board.playerYPos)

	clearConsole()
}

func getDifficultyLevel() string {
	var difficulty string
	validDifficultyLevels := []string{"easy", "medium", "hard"}

	for {
		fmt.Print("Choose a difficulty level (easy, medium, or hard): ")
		fmt.Scanln(&difficulty)
		difficulty = strings.ToLower(difficulty)

		for _, validLevel := range validDifficultyLevels {
			if difficulty == validLevel {
				return difficulty
			}
		}

		fmt.Println("Invalid difficulty level. Please choose from easy, medium, or hard.")
	}
}

func initializeBoard(difficultyLevel string) *board {
	switch difficultyLevel {
	case "hard":
		width, height := 25, 30
		return &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) + 1) * 2, playerYPos: (rand.Intn(height) + 1) * 2}
	case "medium":
		width, height := 15, 18
		return &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) + 1) * 2, playerYPos: (rand.Intn(height) + 1) * 2}
	case "easy":
		width, height := 10, 12
		return &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) + 1) * 2, playerYPos: (rand.Intn(height) + 1) * 2}
	}
	return &board{}
}
