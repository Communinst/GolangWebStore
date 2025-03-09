package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
	entities "github.com/Communinst/GolangWebStore/backend/entity"
	Repository "github.com/Communinst/GolangWebStore/backend/repository"
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

	repository := Repository.New(db)

	repository.UserRepo.PostUser(context.Background(), &entities.User{
		Login:      "moron",
		Password:   "Slaveyan",
		Nickname:   "Gypsy Crusader",
		Email:      "GoodOldCrusafiction@execute.com",
		SignUpDate: time.Now(),
		RoleId:     1,
	})
	repository.UserRepo.PostUser(context.Background(), &entities.User{
		Login:      "gay",
		Password:   "Slaveyan",
		Nickname:   "Gypsy Crusader",
		Email:      "GodOldCrusafiction@execute.com",
		SignUpDate: time.Now(),
		RoleId:     1,
	})
	data, _ := repository.UserRepo.GetUser(context.Background(), 1)
	data.Print()

	//repository.UserRepo.PutUserRole(context.Background(), 1, 2)
	//repository.UserRepo.DeleteUser(context.Background(), 1)

	strg.RunDBTableScript(db, "X:\\Coding\\Golang\\backend\\PostgreSQLScripts\\drop.sql")
	if err := strg.CloseDBConn(db); err != nil {
		log.Fatal("Failed to cease DB connection!")
	} else {
		log.Print("Successfully ceased DB connection!")
	}
}
