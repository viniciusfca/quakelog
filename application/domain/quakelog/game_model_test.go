package quakelog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderKills(t *testing.T) {
	game := Game{
		KillsMap: map[string]int{
			"John":  3,
			"Jane":  5,
			"Steve": 1,
		},
	}

	game.OrderKills()

	expectedKills := []PlayerKill{
		{Name: "Jane", Kills: 5},
		{Name: "John", Kills: 3},
		{Name: "Steve", Kills: 1},
	}

	assert.Equal(t, expectedKills, game.Kills, "The two slices should be the same.")
}

func TestAddPlayerIfNotExists(t *testing.T) {
	game := Game{
		Players: []string{"Jane"},
	}

	game.AddPlayerIfNotExists("John")
	assert.ElementsMatch(t, []string{"Jane", "John"}, game.Players, "The players should match.")

	game.AddPlayerIfNotExists("<world>") // <world> shouldn't be added
	assert.ElementsMatch(t, []string{"Jane", "John"}, game.Players, "The players should match.")

	game.AddPlayerIfNotExists("Jane") // Jane is already present
	assert.ElementsMatch(t, []string{"Jane", "John"}, game.Players, "The players should match.")
}
