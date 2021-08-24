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

	DropNoteTable() error
	CreateNoteTable() error

	CreatePermissionTable() error
	DropPermissionTable() error

	CreateUser(username, password string) (int, error)
	GetUserByID(userID int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetPasswordHash(username string) (string, error)
	SearchForUsername(username string) ([]string, error)

	CreateNote(note models.Note) (int, error)
	GetNoteByID(noteID int) (*models.Note, error)
	GetNotesByOwner(userID int) ([]models.Note, error)
	GetNotesByDelegatedUser(userID int) ([]models.Note, error)
	GetNotesSharedWith(userId int) ([]models.Note, error)

	CreatePermission(noteID, userID int, permission string) error
	RemovePermission(noteID, userID int) error
	UpdatePermission(noteID, userID int, permission string) error
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
