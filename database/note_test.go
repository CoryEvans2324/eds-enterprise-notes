package database

import (
	"testing"
	"time"

	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

func TestNote(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())

	Mgr.DropUserTable()
	Mgr.CreateUserTable()
	Mgr.DropNoteTable()
	Mgr.CreateNoteTable()

	userId, err := Mgr.InsertUser("test", "test")
	owner, _ := Mgr.GetUserByID(userId)
	checkErrNil(t, err)

	dueDate := time.Now().Add(4 * time.Hour)
	note := models.Note{
		Name:    "TestNote",
		Content: "This is a test",
		Owner:   owner,
		DueDate: &dueDate,
	}

	// Without delegated user
	noteID, err := Mgr.InsertNote(note)
	checkErrNil(t, err)

	noteReturned, err := Mgr.GetNoteByID(noteID)
	checkErrNil(t, err)

	if noteReturned.DelegatedUser != nil {
		t.Error("Delegated user should have been nil")
	}

	// With delegated user
	deleID, _ := Mgr.InsertUser("deleUser", "test")

	dele, _ := Mgr.GetUserByID(deleID)

	note.DelegatedUser = dele

	noteID, err = Mgr.InsertNote(note)
	checkErrNil(t, err)

	noteReturned, err = Mgr.GetNoteByID(noteID)
	checkErrNil(t, err)

	if noteReturned.DelegatedUser == nil {
		t.Error("Delegated user should not be nil")
	}

	// Get note that doens't exist
	_, err = Mgr.GetNoteByID(45679)
	if err == nil {
		t.Error("Note with id: 45679 should not exist")
	}

	// Insert and get note where owner or delegatedUser doesn't exist
	note.DelegatedUser = &models.User{
		UserID: 234567890,
	}

	id, err := Mgr.InsertNote(note)
	checkErrNil(t, err)
	_, err = Mgr.GetNoteByID(id)
	if err == nil {
		t.Error("Should have failed when getting note with broken delegatedUserID")
	}

	note.Owner = &models.User{
		UserID: 6782390,
	}

	id, err = Mgr.InsertNote(note)
	checkErrNil(t, err)
	_, err = Mgr.GetNoteByID(id)
	if err == nil {
		t.Error("Should have failed with ownerID that doesn't exist")
	}
}
