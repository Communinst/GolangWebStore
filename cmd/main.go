package main

import (
	"log"
	"log/slog"
	"os"

	
	strg "github.com/Communinst/GolangWebStore/backend/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main() {

	if err := InitEnv(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	config := setupConfig()

	db := strg.InitDBConn(&config.Database)

	if err := strg.CloseDBConn(db); err != nil {
		log.Fatal("Failed to cease DB connection!")
	} else {
		log.Print("Successfully ceased DB connection!")
	}

	return
}

