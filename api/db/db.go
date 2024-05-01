package db

import (
	"database/sql"
	"fmt"
	"money-api-transfer/api/config"

	_ "github.com/lib/pq"
)

// NewAppDB will return DB client
func NewAppDB(dbConf config.Database) (*sql.DB, error) {
	dbConnInfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		dbConf.Url, dbConf.Port, dbConf.Name, dbConf.Username, dbConf.Password, dbConf.SslMode)

	// Open a connection to the database
	db, err := sql.Open("postgres", dbConnInfo)
	if err != nil {
		return nil, err
	}

	// return db instance
	return db, nil
}
