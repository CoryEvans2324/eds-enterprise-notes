package database

import (
	"testing"
	"time"

	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

func TestNote(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())
	ResetDB()

	owner, err := Mgr.CreateUser("test", "test")
	checkErrNil(t, err)

	dueDate := time.Now().Add(4 * time.Hour)
	note := &models.Note{
		Name:    "TestNote",
		Content: "This is a test",
		Owner:   owner,
		DueDate: &dueDate,
	}

	// Without delegated user
	note, err = Mgr.CreateNote(note)
	checkErrNil(t, err)
	t.Log(note, err)

	noteReturned, err := Mgr.GetNoteByID(note.ID)
	checkErrNil(t, err)
	t.Log(noteReturned, err)

	if noteReturned.DelegatedUser != nil {
		t.Error("Delegated user should have been nil")
	}

	// With delegated user
	dele, _ := Mgr.CreateUser("deleUser", "test")

	note.DelegatedUser = dele

	noteNew, err := Mgr.UpdateNote(note)
	checkErrNil(t, err)

	noteReturned, err = Mgr.GetNoteByID(noteNew.ID)
	checkErrNil(t, err)

	if noteReturned.DelegatedUser == nil {
		t.Error("Delegated user should not be nil")
	}

	// Get note that doens't exist
	_, err = Mgr.GetNoteByID(45679)
	if err == nil {
		t.Error("Note with id: 45679 should not exist")
	}
}
