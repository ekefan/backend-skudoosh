package db

import (
	// "context"
	"database/sql"
	// "fmt"
)


type Store interface {
	Querier
}

type DbQuerier struct {
	db *sql.DB
	*Queries
}


func NewStore(db *sql.DB) Store {
	return DbQuerier{
		db: db,
		Queries: New(db),
	}
}


//Every Transaction would occur here