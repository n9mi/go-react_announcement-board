package main

import (
	"log"
	"server/config"
	"server/database"
	"server/router"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	isUsingDotEnv := true
	if isUsingDotEnv {
		godotenv.Load()
	}
	dbConfig := config.GetDBConfig()

	db, errDb := database.NewDB(dbConfig)
	if errDb != nil {
		log.Fatal(errDb.Error())
	}
	defer database.Drop(db)
	database.Drop(db)
	database.Migrate(db)
	database.Seed(db)

	e := router.InitializeEcho()
	router.AssignRouter(e, db, validator.New())

	e.Logger.Fatal(e.Start(":5000"))
}
