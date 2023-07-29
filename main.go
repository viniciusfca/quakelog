package main

import (
	"fmt"
	"net/http"
	"quakelog/application/domain/quakelog"
	"quakelog/application/endpoints"
	"quakelog/application/infra/database"
	"quakelog/application/shared/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client
var service quakelog.GameService

func init() {

	mongoClient = config.MongoDBConnect()
	repo := database.NewRepository(mongoClient, "quake", "game")
	service := quakelog.NewGameService(repo)
	service.DropCollection()
	service.ParseLog("qgames.log")
}

func main() {
	repo := database.NewRepository(mongoClient, "quake", "game")
	service := quakelog.NewGameService(repo)

	handler := endpoints.Handler{
		GameService: *service,
	}

	router := handler.InitializeRoutes()

	fmt.Println("API Started: 3000...")

	http.ListenAndServe(":3000", router)

}
