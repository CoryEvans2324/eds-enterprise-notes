package database

import (
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
	"gorm.io/gorm/clause"
)

func (dbm *databasemanager) CreateNote(note *models.Note) (*models.Note, error) {
	result := dbm.db.Create(&note)
	return note, result.Error
}

func (dbm *databasemanager) UpdateNote(note *models.Note) (*models.Note, error) {
	result := dbm.db.Save(&note)
	return note, result.Error
}

func (dbm *databasemanager) DeleteNote(note *models.Note) error {
	tx := dbm.db.Delete(note)
	return tx.Error
}

func (dbm *databasemanager) GetNoteByID(noteID uint) (*models.Note, error) {
	var note models.Note
	result := dbm.db.Preload("SharedUsers.User").Preload(clause.Associations).First(&note, noteID)

	return &note, result.Error
}

func (dbm *databasemanager) GetNotesByOwner(user *models.User) ([]models.Note, error) {
	var notes = make([]models.Note, 0)
	result := dbm.db.Where("Owner_ID = ?", user.ID).Preload(clause.Associations).Find(&notes)

	return notes, result.Error
}

func (dbm *databasemanager) GetNotesByDelegatedUser(user *models.User) ([]models.Note, error) {
	var notes = make([]models.Note, 0)
	result := dbm.db.Where("Delegated_User_ID = ?", user.ID).Preload(clause.Associations).Find(&notes)

	return notes, result.Error
}

func (dbm *databasemanager) GetNotesSharedWith(user *models.User) ([]models.Note, error) {
	var notes = make([]models.Note, 0)
	result := dbm.db.Model(&models.Note{}).Joins("INNER JOIN Permissions ON Permissions.Note_ID = Notes.ID").Preload(clause.Associations).Find(&notes, "Permissions.User_ID = ?", user.ID)
	return notes, result.Error
}
