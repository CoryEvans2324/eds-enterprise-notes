package database

import (
	"database/sql"

	"github.com/CoryEvans2324/eds-enterprise-notes/models"
	_ "github.com/lib/pq"
)

type DatabaseManager interface {
	Close()
	DropUserTable() error
	CreateUserTable() error

	InsertUser(username, password string) (int, error)
	GetUserByID(userID int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetPasswordHash(username string) (string, error)
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
