package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Storage struct {
		db *sql.DB
	}
)

func SetupStorage() (*Storage, error) {
	dbObj, err := sql.Open("mysql", "root:13BCA006@shiats@tcp(localhost:3306)/candyshop")
	if err != nil {
		return nil, err
	}
	return &Storage{db: dbObj}, err
}
