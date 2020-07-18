package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
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

// ErrInvalidMethod is an error for when the HTTP method is wrong
var ErrInvalidMethod error = errors.New("Invalid method only [get] method is allowed")

// ErrBotIsX is an error for when botIsX is not 1 or 0
var ErrBotIsX error = errors.New("BotIsX should be either 1 or 0")

// ErrLength is an error for when requested state does not have 9 digits
var ErrLength error = errors.New("Length does not match 9 digits")

// ErrStateType is an error for when requested state contains number other than 0, 1, or 2
var ErrStateType error = errors.New("State numbers should be 0, 1, or 2")

// ErrUnmarshalFail is an error for when the Body is unable to be unmarshaled
var ErrUnmarshalFail error = errors.New("Invalid body framing, unable to unmarshal JSON")

// marshal JSON
func constructResponse(res Response) string {
	resp, _ := json.Marshal(res)
	return string(resp[:])
}

// HandleRequest Lambda
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var state [9]int
	var botIsX bool
	var request Body

	if req.HTTPMethod == "GET" { // Allow only GET method
		// Unmarshal Request.body JSON
		err := json.Unmarshal([]byte(req.Body), &request)
		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 400}, ErrUnmarshalFail
		}
		stringArr := strings.Split(request.State, "")
		if request.BotIsX == 0 {
			botIsX = false
		} else if request.BotIsX == 1 {
			botIsX = true
		} else {
			return events.APIGatewayProxyResponse{Body: constructResponse(Response{Action: -1, Status: "Failed", Err: ErrBotIsX})}, ErrBotIsX
		}
		// Length doesn't match 9
		if len(stringArr) != 9 {
			return events.APIGatewayProxyResponse{Body: constructResponse(Response{Action: -1, Status: "Failed", Err: ErrLength})}, ErrLength
		}
		for idx, pos := range stringArr { // Convert to integers
			posInt, err := strconv.Atoi(pos)
			if err != nil { // Cannot parse none number to integer
				return events.APIGatewayProxyResponse{Body: constructResponse(Response{Action: -1, Status: "Failed", Err: err})}, err
			}
			if posInt == 1 || posInt == 2 || posInt == 0 {
				state[idx] = posInt
			} else { // State contains other number than 0, 1, or 2
				return events.APIGatewayProxyResponse{Body: constructResponse(Response{Action: -1, Status: "Failed", Err: ErrStateType})}, ErrStateType
			}
		}
		return events.APIGatewayProxyResponse{Body: constructResponse(Response{Action: ttt.GetNextAIMove(state, botIsX), Status: "Success", Err: nil})}, nil
	}
	return events.APIGatewayProxyResponse{StatusCode: 405}, ErrInvalidMethod
}

func main() {
	lambda.Start(HandleRequest)
}
