package game

import "fmt"

type Board struct {
	Slots [9]Player
}

func NewBoard() Board {
	return Board{Slots: [9]Player{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty}}
}

func (b *Board) Print() {
	for index, value := range b.Slots {
		if value == Empty {
			fmt.Print(index + 1)
		} else {
			fmt.Print(value)
		}

		if index == 2 || index == 5 {
			fmt.Println()
			fmt.Println("——————————")
		} else if index < 2 || index < 5 || index < 8 {
			fmt.Print(" | ")
		}
	}
}

func (b *Board) IsWinner(p Player) bool {
	if p == Empty {
		return false
	}

	slots := b.Slots

	return (slots[0] == p && slots[1] == p && slots[2] == p) ||
		(slots[3] == p && slots[4] == p && slots[5] == p) ||
		(slots[6] == p && slots[7] == p && slots[8] == p) ||
		// diagonals
		(slots[0] == p && slots[4] == p && slots[8] == p) ||
		(slots[2] == p && slots[4] == p && slots[6] == p) ||
		// columns
		(slots[0] == p && slots[3] == p && slots[6] == p) ||
		(slots[1] == p && slots[4] == p && slots[7] == p) ||
		(slots[2] == p && slots[5] == p && slots[8] == p)
}

func (b *Board) InsertPlayerMoveAt(p Player, index int) {
	b.Slots[index] = p
}

func (b *Board) IsSlotEmpty(index int) bool {
	return b.Slots[index] == Empty
}
