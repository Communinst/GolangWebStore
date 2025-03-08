package main

import (
	"log"
	"log/slog"
	"os"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
	strg "github.com/Communinst/GolangWebStore/backend/storage"
	"github.com/Communinst/GolangWebStore/backend/repository"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupConfig() *cnfg.Config {
	config, err := cnfg.LoadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return config
}

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
	return err
}

func main() {

	if err := InitEnv(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	config := setupConfig()

	db := strg.InitDBConn(&config.Database)
	strg.RunDBTableScript(db, "X:\\Coding\\Golang\\backend\\PostgreSQLScripts\\init.sql")

	repository := repository.newRepsitory(db)

	strg.RunDBTableScript(db, "X:\\Coding\\Golang\\backend\\PostgreSQLScripts\\drop.sql")
	if err := strg.CloseDBConn(db); err != nil {
		log.Fatal("Failed to cease DB connection!")
	} else {
		log.Print("Successfully ceased DB connection!")
	}
}
