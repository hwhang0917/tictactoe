package main

import (
	"fmt"

	"github.com/hwhang0917/tictactoe/ttt"
)

func main() {

	var game = [9]int{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0}
	var play = ttt.GetNextAIMove(game, false)
	fmt.Println(play)
}
