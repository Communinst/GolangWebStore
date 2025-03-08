package storage

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"

	cnfg "github.com/Communinst/GolangWebStore/backend/config"
	"github.com/jmoiron/sqlx"
)

func InitDBConn(config *cnfg.Database) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DBName)
	db, err := sqlx.Open("postgres", connStr)
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

func DefineDBTableByScript(db *sqlx.DB, scriptPath string) error {

	script, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to read SQL script file: %v", err)
	}

	tx, err := db.Beginx()
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	_, err = db.ExecContext(ctx, string(script))
	if err != nil {
		log.Fatalf("Failed to execute SQL script: %v", err)
	}

}

func CloseDBConn(db *sql.DB) error {
	return db.Close()
}
