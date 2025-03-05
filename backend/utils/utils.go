package utils

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	cnfg "github.com/Communinst/GolangWebStore/backend/config"
)

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}
	return err
}
func setupConfig() *cnfg.Config {
	config, err := cnfg.LoadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return config
}
