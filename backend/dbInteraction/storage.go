package storage

import (
	"database/sql"
	"fmt"
	"log"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
)

func OpenDB(config *cnfg.Database) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open database connection: ", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
		return nil
	}
	log.Print("Successfully connected to the database!")
	return db
}

func CloseDb(db *sql.DB) error {
	return db.Close()
}
