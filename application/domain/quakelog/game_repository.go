package quakelog

type GameRepository interface {
	Save(games []Game) error
	FindByID(id int) (*Game, error)
	FindAll() ([]Game, error)
	DropCollection() error
}
