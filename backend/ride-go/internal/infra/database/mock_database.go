package database

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMockDatabase() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
