package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const WINNING_SCORE = 100
const GAMES_PER_ROUND = 10

func parseArgs(args []string) []int {
	intArgs := make([]int, 0, len(args))
	for _, arg := range args {
		intArg, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalln(err)
		}
		intArgs = append(intArgs, intArg)
	}

	return intArgs
}

func rollDice() int {
	return 1 + rand.Intn(6)
}

func playTurn(strat int) int {
	score := 0
	for score <= strat {
		rolled := rollDice()
		if rolled == 1 {
			return 0
		}

		score += rolled
	}

	return score
}

func playGame(p1Strat int, p2Strat int) int {
	var p1Score, p2Score int

	for {
		p1Score += playTurn(p1Strat)
		if p1Score >= WINNING_SCORE {
			return 0
		}
		p2Score += playTurn(p2Strat)
		if p2Score >= WINNING_SCORE {
			return 1
		}
	}
}

func calculatePercentage(part, total int) float32 {
	return float32(part) * 100 / float32(total)
}

func playRound(p1Strat int, p2Strat, games int) string {
	p1Wins := 0
	for i := 0; i < games; i++ {
		if playGame(p1Strat, p2Strat) == 0 {
			p1Wins++
		}
	}

	p1Losses := games - p1Wins
	p1WinPercentage := calculatePercentage(p1Wins, games)
	p1LossPercentage := 100 - p1WinPercentage
	return fmt.Sprintf("Holding at %d vs Holding at %d: wins: %d/%d (%0.1f%%), losses: %d/%d(%0.1f%%)", p1Strat, p2Strat, p1Wins, games, p1WinPercentage, p1Losses, games, p1LossPercentage)
}

func main() {
	args := parseArgs(os.Args[1:])
	p1Strat, p2Strat := args[0], args[1]
	fmt.Println(args)

	fmt.Println(playRound(p1Strat, p2Strat, GAMES_PER_ROUND))
}
