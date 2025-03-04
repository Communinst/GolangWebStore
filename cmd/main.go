package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
	strg "github.com/Communinst/GolangWebStore/backend/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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
	fmt.Println("mrn")
	db := strg.OpenDB(&config.Database)

	server := server

	if err := strg.CloseDb(db); err != nil {
		log.Fatal("Failed to cease DB connection!")
	} else {
		log.Print("Successfully ceased DB connection!")
	}

	return
}

func setupConfig() *cnfg.Config {
	config, err := cnfg.LoadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return config
}
