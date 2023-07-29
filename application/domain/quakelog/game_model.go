package quakelog

import "sort"

type PlayerKill struct {
	Name  string `json:"player"`
	Kills int    `json:"kills"`
}

type Game struct {
	ID          int            `json:"id"`
	TotalKills  int            `json:"total_kills"`
	Players     []string       `json:"players"`
	KillsMap    map[string]int `json:"killsMap,omitempty"`
	Kills       []PlayerKill   `json:"kills"`
	KillsByMean map[string]int `json:"kills_by_means"`
}

func (g *Game) OrderKills() {
	var playerKills []PlayerKill
	for name, kills := range g.KillsMap {
		playerKills = append(playerKills, PlayerKill{Name: name, Kills: kills})
	}
	sort.Slice(playerKills, func(i, j int) bool {
		return playerKills[i].Kills > playerKills[j].Kills
	})
	g.Kills = playerKills
}

func (g *Game) AddPlayerIfNotExists(playerName string) {
	for _, player := range g.Players {
		if playerName == "<world>" || player == playerName {
			return
		}
	}

	g.Players = append(g.Players, playerName)
}
