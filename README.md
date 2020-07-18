# TicTacToe

[![Author](https://img.shields.io/badge/author-hwhang0917-green?style=flat)](https://github.com/hwhang0917/tictactoe)
[![License](https://img.shields.io/github/license/hwhang0917/tictactoe)](https://github.com/hwhang0917/tictactoe/blob/master/LICENSE)

> TicTacToe with AI implemented via minimax algorithm

## TODO

- [ ] Make AWS Lambda into HTTP API
- [ ] Build Frontend web with React JS

## Deployment

The function is deployed on AWS Lambda.

## Request

Request with the following body:

```json
{
    "state": "000000000",
    "botIsX": 0 or 1
}
```

#### State

state represents the current board of the tic-tac-toe table.

| Play  | State Number |
| ----- | ------------ |
| O     | 1            |
| X     | 2            |
| Empty | 0            |

For instance the board below would be represented with the following state string:

```json
{
  "state": "012020100"
}
```

![](./assets/012020100.png)

#### BotIsX

botIsX represents whether the box is X or O. If the bot is playing as X value is 1, otherwise the value is 0.

```json
{
  "botIsX": 1
}
```

## Response

Response would look like the following:

```json
{
  "action": 5,
  "status": "Success",
  "error": null
}
```

#### Action

action represents where the bot would play next. It will return -1 if there is an error.

#### Status

status shows rather the request was successful or not.

#### Error

error will show what the exact error is.
The following are type of errors that can happen:

| Errors               | Expected                             |
| -------------------- | ------------------------------------ |
| State Length Error   | State should always be 9 digits long |
| Invalid State Error  | State should either be 0, 1, or 2    |
| Invalid BotIsX Error | BotIsX should always be 0 or 1       |
