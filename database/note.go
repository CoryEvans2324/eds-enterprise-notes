package database

import (
	"fmt"

	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

const NilDelegatedUserValue = -1

func (dbm *databasemanager) InsertNote(note models.Note) (int, error) {
	var delegatedUserID int
	if note.DelegatedUser == nil {
		delegatedUserID = NilDelegatedUserValue
	} else {
		delegatedUserID = note.DelegatedUser.UserID
	}

	row := dbm.db.QueryRow(`
INSERT INTO note (name, content, status, ownerID, dueDate, delegatedUserID)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING noteID`,
		note.Name,
		note.Content,
		note.Status,
		note.Owner.UserID,
		note.DueDate,
		delegatedUserID,
	)

	var noteID int
	err := row.Scan(&noteID)

	return noteID, err
}

func (dbm *databasemanager) GetNoteByID(noteID int) (*models.Note, error) {
	row := dbm.db.QueryRow(`
SELECT name, content, status, ownerID, dueDate, delegatedUserID
FROM note
WHERE noteID = $1`, noteID)

	note := &models.Note{
		NoteID: noteID,
	}
	var ownerID int
	var deleID int

	err := row.Scan(&note.Name, &note.Content, &note.Status, &ownerID, &note.DueDate, &deleID)
	if err != nil {
		return nil, err
	}

	owner, err := dbm.GetUserByID(ownerID)
	if err != nil {
		return nil, err
	}

	note.Owner = owner

	if deleID != NilDelegatedUserValue {
		dele, err := dbm.GetUserByID(deleID)
		if err != nil {
			return nil, fmt.Errorf("delegated User doesn't exist: %v", err)
		}

		note.DelegatedUser = dele
	}

	return note, nil
}
