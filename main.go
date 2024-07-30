package main

import (
	"fmt"
	"log"
	"tictac-toe/game"
)

func main() {
	board := game.NewBoard()
	currentPlayer := game.X

	for {
		board.Print()
		fmt.Printf("\n%s's turn. Select slot: ", currentPlayer)

		var index int
		for {
			_, err := fmt.Scan(&index)
			if err != nil {
				log.Fatal(err)
			}

			if board.IsSlotEmpty(index - 1) {
				break
			} else {
				fmt.Printf("Slot %d has been taken! Pick another slot: ", index)
			}
		}

		board.InsertPlayerMoveAt(currentPlayer, index-1)

		if board.IsWinner(currentPlayer) {
			fmt.Printf("%s won!", currentPlayer)
			break
		} else {
			if currentPlayer == game.X {
				currentPlayer = game.O
			} else {
				currentPlayer = game.X
			}
		}
	}

}
