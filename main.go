package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

const NUMBER_OF_GAMES = 10
const MAX_POINTS = 100

type Player struct {
	hold   int
	points int
	wins   int
}

func (p *Player) play() bool {
	points := 0
	// keep rolling until either 1 rolled or reached hold
	for p.points < MAX_POINTS && points < p.hold {
		val := rollDice()
		if val == 1 {
			points = 0
			break
		}

		points += val
	}

	fmt.Println(p.points)
	p.points += points
	return p.points >= MAX_POINTS
}

// input: x, rolling till total >= x
func parseArgsToInt(args []string) []int {
	intArgs := make([]int, 0, len(args))

	for _, argString := range args {
		num, err := strconv.Atoi(argString)
		if err != nil {
			log.Fatalln("Only integer arguments are allowed", err)
		}

		intArgs = append(intArgs, num)
	}

	return intArgs
}

func rollDice() int {
	return 1 + rand.Intn(6)
}

func startGame(p1, p2 *Player) {
	for {
		// returns true, if reached MAX_points
		won := p1.play()
		if won {
			p1.wins++
			break
		}

		won = p2.play()
		if won {
			p2.wins++
			break
		}
	}

	p1.points = 0
	p2.points = 0
}

func main() {
	intArgs := parseArgsToInt(os.Args[1:])
	fmt.Println(intArgs)

	p1 := Player{hold: intArgs[0]}
	p2 := Player{hold: intArgs[1]}

	for i := 0; i < NUMBER_OF_GAMES; i++ {
		// returns playerId of winner
		startGame(&p1, &p2)
	}

	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)", p1.hold, p2.hold, p1.wins, NUMBER_OF_GAMES, float32(p1.wins*100/NUMBER_OF_GAMES), p2.wins, NUMBER_OF_GAMES, float32(p2.wins*100/NUMBER_OF_GAMES))
}
