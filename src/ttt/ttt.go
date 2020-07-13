package ttt

import (
	"log"
)

// O is the integer value for O player
const O int = 1

// X is the integer value for X player
const X int = 2

// EMPTY is the integer value for empty space
const EMPTY int = 0

// EvaluateGame evaluates the game with given state (1: Max player wins, -1: Min player wins, 0: Draw or Game is still on)
func EvaluateGame(state [9]int) int {
	// Winning cases' indexes
	winningCases := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Row
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Column
		{0, 4, 8}, {2, 4, 6}, // Diagonal
	}

	// O,X Count
	var (
		maxCount int
		minCount int
	)

	for _, winningCase := range winningCases {
		maxCount = 0
		minCount = 0
		for _, winIdx := range winningCase {
			switch state[winIdx] {
			case X:
				maxCount++
				break
			case O:
				minCount++
				break
			case EMPTY:
				break
			default: // Error with state
				log.Fatal("[ERROR] There is something wrong with the state. Values should be 0,1,or 2.")
			}
		}

		// Check for winner
		if maxCount == 3 { // Player 1 wins
			return 1
		} else if minCount == 3 { // Player 2 wins
			return -1
		}
	}
	// Game is still on || Game is a Draw
	return 0
}

// IsFinished returns 1 if the game is finished, 0 if is still on, -1 if error
func IsFinished(state [9]int) bool {
	emptyCell := 9

	for _, box := range state {
		if box == O || box == X {
			emptyCell--
		} else if box == EMPTY {
			// Do nothing
		} else { // Error state
			log.Fatal("[ERROR] There is something wrong with the state. Values should be 0,1,or 2.")
		}
	}

	if emptyCell == 0 { // Game is finished
		return true
	}
	// Game is still on
	return false
}

// GetNextAIMove gets the next best AI move based on Minimax algorithm
func GetNextAIMove(state [9]int) int {
	action := minimax(state, true)

	return action
}

func minimax(state [9]int, maxTurn bool) int {
	action := 0
	if maxTurn { // Max Player's Turn
		maxv := -2000
		// action = 0
		for i := 0; i < 9; i++ {
			if state[i] == EMPTY { // a valid move
				newState := state
				newState[i] = X
				val := minValue(newState, -2000, 2000)
				if val > maxv {
					maxv = val
					action = i
				}
			}
		}
	} else { // Min Player's Turn
		minv := 2000
		// action = 0
		for i := 0; i < 9; i++ {
			if state[i] == EMPTY { // a valid move
				newState := state
				newState[i] = O
				val := maxValue(newState, -2000, 2000)

				if val < minv {
					minv = val
					action = i
				}
			}
		}
	}
	return action
}

func minValue(state [9]int, alpha int, beta int) int {
	// Terminal Test
	evaluation := EvaluateGame(state)
	if evaluation == 1 || evaluation == -1 { // Winning condition
		return EvaluateGame(state) // Returns 1 or -1
	} else if IsFinished(state) { // Draw
		return EvaluateGame(state)
	}

	utilityValue := 2000 // Arbitrary value
	for i := 0; i < 9; i++ {
		if state[i] == EMPTY {
			newState := state
			newState[i] = O

			tmpValue := maxValue(newState, alpha, beta)

			if tmpValue < utilityValue {
				utilityValue = tmpValue
			}
			if utilityValue <= alpha {
				return utilityValue
			}
			if beta > utilityValue {
				beta = utilityValue
			}
		}
	}
	return utilityValue
}

func maxValue(state [9]int, alpha int, beta int) int {
	// Terminal Test
	evaluation := EvaluateGame(state)
	if evaluation == 1 || evaluation == -1 { // Winning condition
		return EvaluateGame(state) // Returns 1 or -1
	} else if IsFinished(state) { // Draw
		return EvaluateGame(state)
	}

	utilityValue := -2000 // Arbitrary value
	for i := 0; i < 9; i++ {
		if state[i] == EMPTY {
			newState := state
			newState[i] = X

			tmpValue := minValue(newState, alpha, beta)

			if tmpValue > utilityValue {
				utilityValue = tmpValue
			}
			if utilityValue >= beta {
				return utilityValue
			}
			if alpha < utilityValue {
				alpha = utilityValue
			}
		}
	}
	return utilityValue
}
