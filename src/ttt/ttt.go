package ttt

// EvaluateGame evaluates the game with given state (-1: Error, 0: Game continues, 1: Player 1 wins, 2: Player 2 wins)
func EvaluateGame(state [9]int) int {
	winningCases := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Row
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Column
		{0, 4, 8}, {2, 4, 6}, // Diagonal
	}

	// O,X Count
	var (
		p1Count int
		p2Count int
	)

	for _, winningCase := range winningCases {
		p1Count = 0
		p2Count = 0
		for _, winIdx := range winningCase {
			switch state[winIdx] {
			case 1:
				p1Count++
				break
			case 2:
				p2Count++
				break
			case 0:
				break
			default: // Error with state
				return -1
			}
		}

		// Check for winner
		if p1Count == 3 { // Player 1 wins
			return 1
		} else if p2Count == 3 { // Player 2 wins
			return 2
		}
	}
	// Game is still on
	return 0
}
