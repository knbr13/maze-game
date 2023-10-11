package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	difficultyLevel := getDifficultyLevel()
	board := initializeBoard(difficultyLevel)
	board.cells.carve(board.playerXPos, board.playerYPos)
	clearConsole()
	initialTime := flag.Int("time", 60, "time limit to reach the gate in seconds")
	flag.Parse()

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
			*initialTime--
			if *initialTime == 0 {
				fmt.Println("you ran out of time :(")
				return
			}
			clearConsole()
			fmt.Printf("%v\nremaining: %v seconds\n", board, *initialTime)
		case r := <-runeReceiverChan:
			key := <-keyReceiverChan
			board.handleMove(r, key)
			clearConsole()
			fmt.Printf("%v\nremaining: %v seconds\n", board, *initialTime)
			if board.checkWin() {
				fmt.Println("you won the game!")
				return
			}
		}
	}
}
