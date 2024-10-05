package main

import (
	"testing"
)

func TestDistributePlayers2(t *testing.T) {
	players := []Player{
		{"Alice", 175},
		{"Bob", 180},
		{"Charlie", 170},
		{"David", 185},
		{"Eve", 172},
		{"Frank", 178},
		{"Grace", 174},
		{"Hank", 181},
		{"Ivy", 169},
		{"Jack", 168},
	}

	teams := distributePlayers(players)

	// Calculate the total height of each team
	teamHeights := calculateTeamHeights(teams)

	// Check if all teams have the same or similar total height
	maxHeight := teamHeights[0]
	minHeight := teamHeights[0]

	for _, height := range teamHeights {
		if height > maxHeight {
			maxHeight = height
		}
		if height < minHeight {
			minHeight = height
		}
	}

	if maxHeight-minHeight > 10 { // Adjust the delta as needed
		println(maxHeight, minHeight)
		t.Errorf("Teams heights are not balanced: %v", teamHeights)
	}
}

func TestDistributePlayers(t *testing.T) {
	players := []Player{
		{"Alice", 175},
		{"Bob", 180},
		{"Charlie", 170},
		{"David", 185},
		{"Eve", 172},
		{"Frank", 178},
		{"Grace", 174},
		{"Hank", 181},
		{"Ivy", 169},
	}

	teams := distributePlayers(players)

	// Calculate the total height of each team
	teamHeights := calculateTeamHeights(teams)

	// Check if all teams have the same or similar total height
	for i := 1; i < len(teamHeights); i++ {
		println(abs(teamHeights[i] - teamHeights[i-1]))
		if abs(teamHeights[i]-teamHeights[i-1]) > 10 { // Adjust the delta as needed
			t.Errorf("Teams heights are not balanced: %v", teamHeights)
		}
	}

	if len(players)%3 != 0 && len(teams) != (len(players)/3)+1 {
		t.Errorf("Expected number of teams: %d, but got %d", (len(players)/3)+1, len(teams))
	}
}
