package database

import (
	"context"
	"quakelog/application/domain/quakelog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository struct {
	collection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client, dbName, collectionName string) *GameRepository {
	if mongoClient == nil {
		return nil
	}
	collection := mongoClient.Database("quake").Collection("game")
	return &GameRepository{
		collection: collection,
	}
}

func (r *GameRepository) Save(games []quakelog.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var interfaceSlice []interface{} = make([]interface{}, len(games))
	for i, d := range games {
		interfaceSlice[i] = d
	}
	_, err := r.collection.InsertMany(ctx, interfaceSlice)
	return err
}

func (r *GameRepository) FindByID(id int) (*quakelog.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	game := &quakelog.Game{}
	err := r.collection.FindOne(ctx, filter).Decode(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (r *GameRepository) FindAll() ([]quakelog.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var games []quakelog.Game
	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}
	return games, nil
}

func (r *GameRepository) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.collection.Drop(ctx)
}
