package main

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tictactoe/src/ttt"
)

// O is the integer value for O player
const O int = 1

// X is the integer value for X player
const X int = 2

// E is the integer value for empty space
const E int = 0

// Body represents data for TTT function to computer minimax algorithm
type Body struct {
	State  string `json:"state"`
	BotIsX int    `json:"botIsX"`
}

// Response represents data to be sent back to the client
type Response struct {
	Action int    `json:"action"`
	Status string `json:"status"`
	Err    error  `json:"error"`
}

// ErrBotIsX is an error for when botIsX is not 1 or 0
var ErrBotIsX error = errors.New("BotIsX should be either 1 or 0")

// ErrLength is an error for when requested state does not have 9 digits
var ErrLength error = errors.New("Length does not match 9 digits")

// ErrStateType is an error for when requested state contains number other than 0, 1, or 2
var ErrStateType error = errors.New("State numbers should be 0, 1, or 2")

// HandleRequest Lambda
func HandleRequest(ctx context.Context, req Body) (Response, error) {
	var state [9]int
	var botIsX bool
	stringArr := strings.Split(req.State, "")
	if req.BotIsX == 0 {
		botIsX = false
	} else if req.BotIsX == 1 {
		botIsX = true
	} else {
		return Response{Action: -1, Status: "Failed", Err: ErrBotIsX}, ErrBotIsX
	}
	// Length doesn't match 9
	if len(stringArr) != 9 {
		return Response{Action: -1, Status: "Failed", Err: ErrLength}, ErrLength
	}
	for idx, pos := range stringArr { // Convert to integers
		posInt, err := strconv.Atoi(pos)
		if err != nil { // Cannot parse none number to integer
			return Response{Action: -1, Status: "Failed", Err: err}, err
		}
		if posInt == 1 || posInt == 2 || posInt == 0 {
			state[idx] = posInt
		} else { // State contains other number than 0, 1, or 2
			return Response{Action: -1, Status: "Failed", Err: ErrStateType}, ErrStateType
		}
	}
	return Response{Action: ttt.GetNextAIMove(state, botIsX), Status: "Success", Err: nil}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
