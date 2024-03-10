package main

import "testing"

func TestParseArgs(t *testing.T) {
	args := []string{"1", "2", "3"}
	expected := []int{1, 2, 3}

	result := parseArgs(args)

	if len(result) != len(expected) {
		t.Errorf("Expected %d arguments, but got %d", len(expected), len(result))
	}

	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Expected %d at index %d, but got %d", expected[i], i, val)
		}
	}
}

func TestPlayTurn(t *testing.T) {
	strat := 10
	for i := 0; i < 10000; i++ {
		score := playTurn(strat)
		if score > strat+6 {
			t.Errorf("Expected score < %d, but got %d", strat, score)
		}
	}
}
