package database

import (
	"testing"

	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

func TestPermissions(t *testing.T) {
	cfg := createConfig()
	CreateDatabaseManager(cfg.Database.DataSourceName())

	Mgr.DropNoteTable()
	Mgr.CreateNoteTable()

	Mgr.DropUserTable()
	Mgr.CreateUserTable()

	Mgr.DropPermissionTable()
	Mgr.CreatePermissionTable()

	// create test users
	ownerID, _ := Mgr.CreateUser("owner", "password")
	otherID, _ := Mgr.CreateUser("other", "password")

	// create test note
	note := models.Note{
		Name:    "testing",
		Content: "test",
		Owner: &models.User{
			UserID: ownerID,
		},
		Status: "In Progress",
		SharedUsers: []models.Permission{
			{
				Permission: "viewer",
				User: models.User{
					UserID: otherID,
				},
			},
		},
	}

	noteID, err := Mgr.CreateNote(note)
	if err != nil {
		t.Errorf("Cannot create note: %v", err)
	}

	returnedNote, err := Mgr.GetNoteByID(noteID)
	if err != nil {
		t.Errorf("Cannot get note: %v", err)
	}

	// test permissions
	if len(returnedNote.SharedUsers) == 0 {
		t.Error("Note SharedUsers list is empty")
	}

	// remove the permission
	err = Mgr.RemovePermission(noteID, otherID)
	checkErrNil(t, err)

	// create
	err = Mgr.CreatePermission(noteID, otherID, "editor")
	checkErrNil(t, err)

	// update
	err = Mgr.UpdatePermission(noteID, otherID, "viewer")
	checkErrNil(t, err)

}
