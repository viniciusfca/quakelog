package quakelog

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"quakelog/application/errors"
	"regexp"
	"strings"
)

type GameService struct {
	gameRepository GameRepository
}

func NewGameService(repo GameRepository) *GameService {
	return &GameService{
		gameRepository: repo,
	}
}

func (s *GameService) FindById(gameId int) (Game, *errors.APIError) {
	game, err := s.gameRepository.FindByID(gameId)

	if err != nil {
		log.Println("Error to get report for game ID: ", err)

		if err.Error() == "mongo: no documents in result" {
			return Game{}, &errors.APIError{StatusCode: 404, Message: fmt.Sprintf("Report not found for game ID : %d", gameId)}
		}
		return Game{}, &errors.APIError{StatusCode: 500, Message: fmt.Sprintf("Error to get report for game ID : %d", gameId)}
	}
	return *game, nil
}

func (s *GameService) FindAll() ([]Game, *errors.APIError) {
	games, err := s.gameRepository.FindAll()
	if err != nil {
		log.Println("Error to get all games ", err)

		if err.Error() == "mongo: no documents in result" {
			return nil, &errors.APIError{StatusCode: 404, Message: fmt.Sprintf("Report not found for games")}
		}
		return nil, &errors.APIError{StatusCode: 500, Message: "Error to get all games"}
	}
	return games, nil
}

func (s *GameService) DropCollection() error {
	return s.gameRepository.DropCollection()
}

func (s *GameService) ParseLog(fileName string) ([]Game, error) {
	file, err := os.Open(fileName)

	if err != nil {
		panic("Error to read file")
	}

	defer file.Close()

	var games []Game
	var currentGame *Game
	scanner := bufio.NewScanner(file)
	killRegEx := regexp.MustCompile(`Kill: (\d+) (\d+) (\d+): (.+) killed (.+) by (.+)`)
	gameID := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "InitGame:") {
			if currentGame != nil {
				// Ordenando as kills antes de adicionar ao slice de games
				currentGame.OrderKills()
				games = append(games, *currentGame)
			}
			gameID++
			currentGame = &Game{
				ID:          gameID,
				KillsMap:    make(map[string]int),
				Kills:       []PlayerKill{},
				KillsByMean: make(map[string]int),
			}
		} else if strings.Contains(line, "Kill:") {
			currentGame.TotalKills++

			parts := killRegEx.FindStringSubmatch(line)
			if parts != nil {
				killer, victim, cause := parts[4], parts[5], parts[6]

				if killer != "<world>" {
					currentGame.KillsMap[killer]++
				}

				if killer == "<world>" {
					currentGame.KillsMap[victim]--
				}

				currentGame.KillsByMean[cause]++
				currentGame.AddPlayerIfNotExists(victim)
				currentGame.AddPlayerIfNotExists(killer)
			}
		}
	}

	if currentGame != nil {
		currentGame.OrderKills()
		games = append(games, *currentGame)
	}

	err = s.gameRepository.Save(games)

	if err != nil {
		log.Println("Fail to save games in mongoDB:", err)
	}

	return games, err
}
