package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	difficultyLevel := getDifficultyLevel()
	board := initializeBoard(difficultyLevel)
	board.cells.carve(board.playerXPos, board.playerYPos)
	clearConsole()
	initialTime := 60
	fmt.Println(board)

	reader := bufio.NewReader(os.Stdin)
	runeReceiverChan := make(chan rune)
	ticker := time.NewTicker(time.Second)

	for {
		go func() {
			r, _, err := reader.ReadRune()
			if err == nil {
				runeReceiverChan <- r
			}
		}()

		select {
		case <-ticker.C:
			initialTime--
			clearConsole()
			fmt.Printf("%v\nremaining: %v\n", board, initialTime)
		case r := <-runeReceiverChan:
			board.handleMove(r)
			clearConsole()
			fmt.Printf("%v\nremaining: %v\n", board, initialTime)
		}
	}
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
		width, height := 32, 16
		return &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) * 2) + 1, playerYPos: (rand.Intn(height) * 2) + 1}
	case "medium":
		width, height := 24, 10
		return &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) * 2) + 1, playerYPos: (rand.Intn(height) * 2) + 1}
	case "easy":
		width, height := 12, 6
		return &board{cells: initializeCells(width, height), playerXPos: (rand.Intn(width) * 2) + 1, playerYPos: (rand.Intn(height) * 2) + 1}
	}
	return &board{}
}
