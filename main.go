package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Player struct to hold player information
type Player struct {
	Name   string
	Height int
}

const TeamSize = 6

// readCSV function to read player data from CSV file
func readCSV(filename string) ([]Player, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var players []Player
	for _, record := range records {
		height, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		player := Player{
			Name:   record[0],
			Height: height,
		}
		players = append(players, player)
	}
	return players, nil
}

// distributePlayers function to distribute players into balanced teams
// distributePlayers function to distribute players into balanced teams
func distributePlayers(players []Player) [][]Player {
	// Sort players by height in descending order
	sort.Slice(players, func(i, j int) bool {
		return players[i].Height > players[j].Height
	})

	numTeams := len(players) / TeamSize
	if len(players)%TeamSize != 0 {
		numTeams++
	}

	teams := make([][]Player, numTeams)
	teamHeights := make([]int, numTeams)

	for _, player := range players {
		// Find the team with the smallest total height
		minHeightIndex := 0
		for j := 1; j < numTeams; j++ {
			if teamHeights[j] < teamHeights[minHeightIndex] {
				minHeightIndex = j
			}
		}
		// Add the player to that team
		teams[minHeightIndex] = append(teams[minHeightIndex], player)
		teamHeights[minHeightIndex] += player.Height
	}

	// Evaluate and balance teams
	evaluateAndBalanceTeams(teams, teamHeights)

	return teams
}

// evaluateAndBalanceTeams function to re-evaluate and balance the teams by swapping players
func evaluateAndBalanceTeams(teams [][]Player, teamHeights []int) {
	numTeams := len(teams)
	for {
		swapped := false
		for i := 0; i < numTeams; i++ {
			for j := i + 1; j < numTeams; j++ {
				// Try to swap each player in team i with each player in team j
				for m := 0; m < len(teams[i]); m++ {
					for n := 0; n < len(teams[j]); n++ {
						// Calculate new heights if players were swapped
						newHeightI := teamHeights[i] - teams[i][m].Height + teams[j][n].Height
						newHeightJ := teamHeights[j] - teams[j][n].Height + teams[i][m].Height
						if abs(newHeightI-newHeightJ) < abs(teamHeights[i]-teamHeights[j]) {
							// Perform the swap
							teams[i][m], teams[j][n] = teams[j][n], teams[i][m]
							teamHeights[i], teamHeights[j] = newHeightI, newHeightJ
							swapped = true
						}
					}
				}
			}
		}
		if !swapped {
			break
		}
	}
}

// calculateTeamHeights function to calculate the total height of each team
func calculateTeamHeights(teams [][]Player) []int {
	teamHeights := make([]int, len(teams))
	for i, team := range teams {
		for _, player := range team {
			teamHeights[i] += player.Height
		}
	}
	return teamHeights
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// writeTeamsToCSV function to write the teams into a CSV file
func writeTeamsToCSV(filename string, teams [][]Player) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Team", "Name", "age"})

	// Write team data
	for i, team := range teams {
		for _, player := range team {
			writer.Write([]string{
				fmt.Sprintf("Team %d", i+1),
				player.Name,
				fmt.Sprintf("%d", player.Height),
			})
		}
	}

	return nil
}

func main() {
	players, err := readCSV("players.csv")
	if err != nil {
		log.Fatalf("Failed to read CSV: %v", err)
	}

	teams := distributePlayers(players)
	teamHeights := calculateTeamHeights(teams)

	for i, team := range teams {
		fmt.Printf("Team %d:\n", i+1)
		fmt.Printf(" Team Height: %d \n", teamHeights[i])
		for no, player := range team {
			fmt.Printf(" %d %s (%d cm)\n", no+1, player.Name, player.Height)
		}
	}

	// Check if all teams have the same or similar total height
	for i := 1; i < len(teamHeights); i++ {
		delta := abs(teamHeights[i] - teamHeights[i-1])
		fmt.Printf(" Team Height: %d with delta: %d \n", teamHeights[i], delta)
		if abs(teamHeights[i]-teamHeights[i-1]) > 10 { // Adjust the delta as needed
			fmt.Printf("Teams %d age are not balanced: %v \n", i, teamHeights[i])
		}
	}

	// Write the result to CSV
	if err := writeTeamsToCSV("teams.csv", teams); err != nil {
		log.Fatalf("Failed to write teams to CSV: %v", err)
	}

	fmt.Println("Teams have been written to teams.csv")
}
