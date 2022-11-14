package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/src/handler"
	"live-easy-backend/src/repository"
	"live-easy-backend/src/usecase"
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
	oauth := infrastructure.InitOAuth()

	// init Storage
	storage := infrastructure.InitStorage(
		os.Getenv("STORAGE_BASE_URL"),
		os.Getenv("STORAGE_BUCKET_NAME"),
		os.Getenv("STORAGE_FOLDER_NAME"),
	)

	// run migration
	if os.Getenv("DB_USERNAME") == "root" {
		migration := sql.Migration{Db: db.DB}
		migration.RunMigration()
	}

	// init repository
	repo := repository.Init(*db, *oauth, *storage)

	// init usecase
	uc := usecase.Init(repo)

	// init handler
	router := gin.Default()
	rest := handler.Init(router, uc, oauth)
	rest.Run()
}
