package database

import (
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseManager interface {
	AutoMigrate()

	CreateUser(username, password string) (*models.User, error)
	GetUserByID(userID uint) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	SearchForUsername(username string) ([]string, error)

	CreateNote(note *models.Note) (*models.Note, error)
	UpdateNote(note *models.Note) (*models.Note, error)
	DeleteNote(note *models.Note) error
	GetNoteByID(noteID uint) (*models.Note, error)
	GetNotesByOwner(user *models.User) ([]models.Note, error)
	GetNotesByDelegatedUser(user *models.User) ([]models.Note, error)
	GetNotesSharedWith(user *models.User) ([]models.Note, error)

	CreatePermission(models.Permission) error
	RemovePermission(models.Permission) error
	UpdatePermission(models.Permission) error
}

type databasemanager struct {
	db *gorm.DB
}

var Mgr DatabaseManager
var db *gorm.DB

func CreateDatabaseManager(dataSourceName string) error {
	newDB, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return err
	}

	db = newDB

	Mgr = &databasemanager{
		db: newDB,
	}
	return nil
}
