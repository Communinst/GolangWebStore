package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/Communinst/GolangWebStore/backend/repository"
	"github.com/Communinst/GolangWebStore/backend/service"
	strg "github.com/Communinst/GolangWebStore/backend/storage"
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
	strg.RunDBTableScript(db, "X:\\Coding\\Golang\\backend\\PostgreSQLScripts\\roles.sql")

	repository := repository.New(db)
	service := service.New(repository)

	repository.CompanyRepo.PostCompany(context.Background(), &entities.Company{
		Name: "Slaveynia",
	})
	repository.CompanyRepo.PostCompany(context.Background(), &entities.Company{
		Name: "Bastardsk",
	})
	repository.CompanyRepo.GetCompany(context.Background(), 2)
	repository.CompanyRepo.GetAllCompanies(context.Background())

	repository.GameRepo.PostGame(context.Background(), &entities.Game{
		PublisherId: 1,
		Name:        "Slaughtery",
	})
	repository.GameRepo.PostGame(context.Background(), &entities.Game{
		PublisherId: 2,
		Name:        "BadAssovsk",
	})
	repository.GameRepo.GetGame(context.Background(), 2)
	repository.GameRepo.GetAllGames(context.Background())
	repository.GameRepo.DeleteGame(context.Background(), 1)
	repository.CompanyRepo.DeleteCompany(context.Background(), 1)

	strg.RunDBTableScript(db, "X:\\Coding\\Golang\\backend\\PostgreSQLScripts\\drop.sql")
	if err := strg.CloseDBConn(db); err != nil {
		log.Fatal("Failed to cease DB connection!")
	} else {
		log.Print("Successfully ceased DB connection!")
	}
}
