package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DatabaseManager interface {
	Close()
	DropUserTable() error
	CreateUserTable() error
}

type databasemanager struct {
	db *sql.DB
}

var Mgr DatabaseManager

func CreateDatabaseManager(dataSourceName string) error {
	// sql.Open only returns an error on an unkown driver name
	db, _ := sql.Open("postgres", dataSourceName)

	if err := db.Ping(); err != nil {
		return err
	}

	Mgr = &databasemanager{
		db: db,
	}

	return nil
}

func (dbm *databasemanager) Close() {
	dbm.db.Close()
}
