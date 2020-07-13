package main

import (
	"fmt"

	"github.com/tictactoe/src/ttt"
)

// O is the integer value for O player
const O int = 1

// X is the integer value for X player
const X int = 2

// E is the integer value for empty space
const E int = 0

func main() {
	sampleState := [9]int{
		E, X, O,
		E, O, E,
		E, E, E,
	}

	// fmt.Println(ttt.EvaluateGame(sampleState))
	// fmt.Println(ttt.IsFinished(sampleState))

	// Expected 6
	fmt.Println(ttt.GetNextAIMove(sampleState))
}
