package main

import (
	"slices"
	"testing"
)

func TestPlayerPlay(t *testing.T) {
	p := Player{hold: 20, points: 0}
	won := p.play()

	if won {
		t.Errorf("expected false, got true")
	}

	if p.points < 0 || p.points > MAX_POINTS {
		t.Errorf("invalid points value: %d", p.points)
	}
}

func TestParseArgsToInt(t *testing.T) {
	args := []string{"10", "20", "30"}
	expected := []int{10, 20, 30}

	result := parseArgsToInt(args)

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, expected %v", result, expected)
	}
}


func TestRollDice(t *testing.T) {
	result := rollDice()

	if result < 1 || result > 6 {
		t.Errorf("invalid dice roll value: %d", result)
	}
}

func TestStartGame(t *testing.T) {
	p1 := Player{hold: 20, points: 0}
	p2 := Player{hold: 20, points: 0}

	startGame(&p1, &p2)

	if p1.points < 0 || p1.points > MAX_POINTS {
		t.Errorf("invalid points value for player 1: %d", p1.points)
	}

	if p2.points < 0 || p2.points > MAX_POINTS {
		t.Errorf("invalid points value for player 2: %d", p2.points)
	}
}

func TestMain(t *testing.T) {
	// TODO: Write tests for the main function
}
