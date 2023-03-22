/**
 * Copyright 2023 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */

package main

type Move int64

const (
	PAPER Move = iota
	ROCK
	SCISSORS
)

type Outcome int64

const (
	PLAYER_ONE_WINS Outcome = iota
	PLAYER_ONE_LOSES
	TIE
)

type game struct {
	playerOneMove Move
	playerTwoMove Move
	outCome       Outcome
}

var games = []game{
	{playerOneMove: SCISSORS, playerTwoMove: PAPER, outCome: PLAYER_ONE_WINS},
	{playerOneMove: SCISSORS, playerTwoMove: ROCK, outCome: PLAYER_ONE_LOSES},
	{playerOneMove: ROCK, playerTwoMove: SCISSORS, outCome: PLAYER_ONE_WINS},
	{playerOneMove: ROCK, playerTwoMove: PAPER, outCome: PLAYER_ONE_LOSES},
	{playerOneMove: PAPER, playerTwoMove: SCISSORS, outCome: PLAYER_ONE_LOSES},
	{playerOneMove: PAPER, playerTwoMove: ROCK, outCome: PLAYER_ONE_WINS},
}

func Kata(playerOneMove, playerTwoMove Move) Outcome {

	for game := range games {
		if games[game].playerOneMove == playerOneMove &&
			games[game].playerTwoMove == playerTwoMove {

			return games[game].outCome
		}
	}

	return TIE

	// if p1Move == p2Move {
	// 	return TIE
	// }

	// if p1Move == PAPER && p2Move == ROCK {
	// 	return PLAYER_ONE_WINS
	// }

	// if p1Move == PAPER && p2Move == SCISSORS {
	// 	return PLAYER_ONE_LOSES
	// }

	// if p1Move == ROCK && p2Move == SCISSORS {
	// 	return PLAYER_ONE_WINS
	// }

	// if p1Move == ROCK && p2Move == PAPER {
	// 	return PLAYER_ONE_LOSES
	// }

	// if p1Move == SCISSORS && p2Move == PAPER {
	// 	return PLAYER_ONE_WINS
	// }
	
	// if p1Move == SCISSORS && p2Move == ROCK {
	// 	return PLAYER_ONE_LOSES
	// }

	// return PLAYER_ONE_WINS
}
