package handler

import (
	ent "github.com/Communinst/golangWebStore/backend/entities/user.go"
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func (u UserRepository) CreateUser(c context.Context){ 
	ent
}

