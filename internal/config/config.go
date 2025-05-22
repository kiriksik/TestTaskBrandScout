package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/kiriksik/TestTaskBrandScout/internal/database"
)

type ApiConfig struct {
	Queries database.QuotesQueries
}

func InitializeApiConfig() *ApiConfig {
	apiCfg := &ApiConfig{
		Queries: initializeDBQueries(),
	}
	return apiCfg
}

func initializeDBQueries() *database.Queries {
	dbURL := os.Getenv("GOOSE_DBSTRING")
	db, err := sql.Open("postgres", fmt.Sprintf("%s?sslmode=disable", dbURL))
	if err != nil {
		log.Printf("error in db connection: %s", err)
		return nil
	}
	dbQueries := database.New(db)
	log.Println("database connected")
	return dbQueries
}
