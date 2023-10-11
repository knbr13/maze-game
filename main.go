package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	difficultyLevel := getDifficultyLevel()
	board := initializeBoard(difficultyLevel)
	board.cells.carve(board.playerXPos, board.playerYPos)
	clearConsole()
	initialTime := 60
	fmt.Println(board)
	if err := keyboard.Open(); err != nil {
		log.Fatal("error: cannot open keyboard: ", err)
	}

	defer keyboard.Close()

	runeReceiverChan := make(chan rune, 1)
	keyReceiverChan := make(chan keyboard.Key, 1)
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			r, key, err := keyboard.GetSingleKey()
			if err == nil {
				runeReceiverChan <- r
				keyReceiverChan <- key
			}
		}
	}()

	for {
		select {
		case <-ticker.C:
			initialTime--
			if initialTime == 0 {
				fmt.Println("you ran out of time :(")
				return
			}
			clearConsole()
			fmt.Printf("%v\nremaining: %v\n", board, initialTime)
		case r := <-runeReceiverChan:
			key := <-keyReceiverChan
			board.handleMove(r, key)
			clearConsole()
			fmt.Printf("%v\nremaining: %v\n", board, initialTime)
			if board.checkWin() {
				fmt.Println("you won the game!")
				return
			}
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
