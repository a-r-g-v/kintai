package store

import (
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"

	"log"
)

type DB struct {
	conn *sqlx.DB
}

func NewDB(database string) *DB {

	db, err := sqlx.Connect("sqlite3", database)
	if err != nil {
		log.Fatalln(err)
	}

	return &DB{conn: db}
}
