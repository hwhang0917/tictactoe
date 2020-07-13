package main

import (
	"fmt"

	"github.com/tictactoe/src/ttt"
)

func main() {
	sampleState := [9]int{
		2, 2, 1,
		0, 1, 1,
		2, 0, 1,
	}
	fmt.Println(ttt.EvaluateGame(sampleState))
}
