package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
	"github.com/Communinst/GolangWebStore/backend/handler"
	"github.com/Communinst/GolangWebStore/backend/repository"
	"github.com/Communinst/GolangWebStore/backend/server"
	"github.com/Communinst/GolangWebStore/backend/service"
	strg "github.com/Communinst/GolangWebStore/backend/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	adminRole = 6
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
	//gin.SetMode(gin.ReleaseMode)
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
	handler := handler.New(service)

	fmt.Printf("%s", config.Address)

	server := server.New(
		config.Address,
		handler.InitRoutes(),
		config.Timeout,
		config.Timeout,
	)

	server.Run()

	strg.RunDBTableScript(db, "X:\\Coding\\Golang\\backend\\PostgreSQLScripts\\drop.sql")
	if err := strg.CloseDBConn(db); err != nil {
		log.Fatal("Failed to cease DB connection!")
	} else {
		//log.Print("Successfully ceased DB connection!")
	}

}
