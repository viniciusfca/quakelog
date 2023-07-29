package contract

import "quakelog/application/domain/quakelog"

type PlayerKillOut struct {
	Name  string `json:"player"`
	Kills int    `json:"total"`
}

type GameOut struct {
	ID          int             `json:"id"`
	TotalKills  int             `json:"total_kills"`
	Players     []string        `json:"players"`
	Kills       []PlayerKillOut `json:"kills"`
	KillsByMean map[string]int  `json:"kills_by_means"`
}

func ConvertToGameOut(game *quakelog.Game) *GameOut {
	var playerKillsOut []PlayerKillOut
	for _, kill := range game.Kills {
		playerKillsOut = append(playerKillsOut, PlayerKillOut{Name: kill.Name, Kills: kill.Kills})
	}

	return &GameOut{
		ID:          game.ID,
		TotalKills:  game.TotalKills,
		Players:     game.Players,
		Kills:       playerKillsOut,
		KillsByMean: game.KillsByMean,
	}
}

func ConvertGameListToGameOutList(games []quakelog.Game) []GameOut {
	var gamesOut []GameOut
	for _, game := range games {
		gamesOut = append(gamesOut, *ConvertToGameOut(&game))
	}

	return gamesOut
}
