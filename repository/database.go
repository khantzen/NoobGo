package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"../model/db"
)

type Repository interface {
	FindUserByEmail(email string) (*db.UserDb, error)
}

type DB struct {
	*sql.DB
}

func InitDatabase() (*DB, error) {
	db, err := sql.Open("mysql", "root:root@/noobGoDb")

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
