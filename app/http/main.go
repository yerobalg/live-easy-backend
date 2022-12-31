package main

import (
	"os"

	"github.com/joho/godotenv"
	"live-easy-backend/database/sql"
	"live-easy-backend/infrastructure"
	"live-easy-backend/sdk/log"
	"live-easy-backend/src/handler"
	"live-easy-backend/src/repository"
	"live-easy-backend/src/usecase"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	
	// init Logger
	logger := log.Init()

	// init DB
	db, err := sql.Init(logger)
	if err != nil {
		panic(err)
	}

	// init Firebase
	firebase := infrastructure.InitFirebase()

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
	repo := repository.Init(*db, *firebase, *storage)

	// init usecase
	uc := usecase.Init(repo)

	// init handler
	rest := handler.Init(uc, logger)
	rest.Run()
}
