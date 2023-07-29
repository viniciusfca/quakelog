package quakelog

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedGameRepository struct {
	mock.Mock
}

func (m *MockedGameRepository) FindByID(gameId int) (*Game, error) {
	args := m.Called(gameId)
	return args.Get(0).(*Game), args.Error(1)
}

func (m *MockedGameRepository) FindAll() ([]Game, error) {
	args := m.Called()
	return args.Get(0).([]Game), args.Error(1)
}

func (m *MockedGameRepository) DropCollection() error {
	return m.Called().Error(0)
}

func (m *MockedGameRepository) Save(games []Game) error {
	return m.Called(games).Error(0)
}

func TestFindById(t *testing.T) {
	repo := new(MockedGameRepository)
	svc := NewGameService(repo)
	expectedGame := Game{ID: 1}
	repo.On("FindByID", 1).Return(&expectedGame, nil)

	game, err := svc.FindById(1)

	assert.Nil(t, err)
	assert.Equal(t, expectedGame, game)
}

func TestFindByIdWithError(t *testing.T) {
	repo := new(MockedGameRepository)
	svc := NewGameService(repo)
	repo.On("FindByID", 1).Return(&Game{}, errors.New("Report not found for game ID"))

	_, err := svc.FindById(1)

	assert.Equal(t, err.StatusCode, 500)
	assert.NotNil(t, err)
}

func TestFindByIdWithError404(t *testing.T) {
	repo := new(MockedGameRepository)
	svc := NewGameService(repo)
	repo.On("FindByID", 1).Return(&Game{}, errors.New("mongo: no documents in result"))

	_, err := svc.FindById(1)

	assert.Equal(t, err.StatusCode, 404)
	assert.NotNil(t, err)
}

func TestFindAll(t *testing.T) {
	repo := new(MockedGameRepository)
	svc := NewGameService(repo)
	expectedGames := []Game{{ID: 1}, {ID: 2}}
	repo.On("FindAll").Return(expectedGames, nil)

	games, err := svc.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedGames, games)
}

func TestFindAllWithError(t *testing.T) {
	repo := new(MockedGameRepository)
	svc := NewGameService(repo)
	repo.On("FindAll").Return([]Game{}, errors.New("Error to get all games"))

	_, err := svc.FindAll()

	assert.NotNil(t, err)
}

func TestFindAllWithError404(t *testing.T) {
	repo := new(MockedGameRepository)
	svc := NewGameService(repo)
	repo.On("FindAll").Return([]Game{}, errors.New("mongo: no documents in result"))

	_, err := svc.FindAll()

	assert.NotNil(t, err)
}

func TestParseLog(t *testing.T) {
	// 1. Create a sample log file
	tempFile, err := ioutil.TempFile("", "testlog")
	assert.NoError(t, err, "Unexpected error creating temp file.")
	defer os.Remove(tempFile.Name())

	logContent := `
InitGame:
Kill: 2 3 4: John killed Jane by GUN
InitGame:
Kill: 2 3 4: Jane killed John by KNIFE
InitGame:
Kill: 2 3 4: <world> killed John by KNIFE
`
	_, err = tempFile.Write([]byte(logContent))
	assert.NoError(t, err, "Unexpected error writing to temp file.")
	tempFile.Close()

	repo := new(MockedGameRepository)
	svc := NewGameService(repo)

	repo.On("Save", mock.Anything).Return(nil)
	games, err := svc.ParseLog(tempFile.Name())
	assert.NoError(t, err, "Unexpected error during ParseLog.")
	assert.Len(t, games, 3, "Expected to parse 3 games from the log.")

	firstGame := games[0]
	assert.Equal(t, 1, firstGame.ID, "Unexpected game ID for the first game.")
	assert.Equal(t, 1, firstGame.TotalKills, "Unexpected total kills for the first game.")
}

func TestParseLogSaveError(t *testing.T) {
	// 1. Create a sample log file
	tempFile, err := ioutil.TempFile("", "testlog")
	assert.NoError(t, err, "Unexpected error creating temp file.")
	defer os.Remove(tempFile.Name())

	logContent := `
InitGame:
Kill: 2 3 4: John killed Jane by GUN
InitGame:
Kill: 2 3 4: Jane killed John by KNIFE
InitGame:
Kill: 2 3 4: <world> killed John by KNIFE
`
	_, err = tempFile.Write([]byte(logContent))
	assert.NoError(t, err, "Unexpected error writing to temp file.")
	tempFile.Close()

	repo := new(MockedGameRepository)
	svc := NewGameService(repo)

	repo.On("Save", mock.Anything).Return(errors.New("Error to save on mongodb"))
	_, err = svc.ParseLog(tempFile.Name())

	assert.NotNil(t, err)
}
