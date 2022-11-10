package main

import (
	"os"

	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/src/handler"
	"live-easy-backend/src/repository"
	"live-easy-backend/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// init DB
	db, err := sql.InitDB()
	if err != nil {
		panic(err)
	}

	// init OAuth
	oauth := infrastructure.OAuth{}
	oauth.Init()

	// init Storage
	storage := infrastructure.Storage{}
	storage.Init("live-easy-bucket")

	// run migration
	if os.Getenv("DB_USERNAME") == "root" {
		migration := sql.Migration{Db: db.DB}
		migration.RunMigration()
	}

	// init repository
	repo := repository.Init(*db, oauth)

	// init usecase
	uc := usecase.Init(repo)

	// init handler
	router := gin.Default()
	rest := handler.Init(router, uc)
	rest.Run()
}
