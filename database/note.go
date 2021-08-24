package database

import (
	"database/sql"
	"fmt"

	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

const NilDelegatedUserValue = -1

func (dbm *databasemanager) CreateNote(note models.Note) (int, error) {
	var delegatedUserID int
	if note.DelegatedUser == nil {
		delegatedUserID = NilDelegatedUserValue
	} else {
		delegatedUserID = note.DelegatedUser.UserID
	}

	tx, err := dbm.db.Begin()
	if err != nil {
		return 0, err
	}
	row := tx.QueryRow(`
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
	err = row.Scan(&noteID)
	if err != nil {
		// rollback and return
		tx.Rollback()
		return noteID, err
	}

	stmt, _ := tx.Prepare(`INSERT INTO notePermission (noteID, userID, permission) VALUES ($1, $2, $3)`)
	for _, permission := range note.SharedUsers {
		_, err := stmt.Exec(noteID, permission.User.UserID, permission.Permission)
		if err != nil {
			tx.Rollback()
			return noteID, err
		}
	}

	err = tx.Commit()

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

	// get the list of permissions
	rows, err := dbm.db.Query(`
SELECT noteUser.UserID, noteUser.Username, noteUser.userRole, notePermission.permission from note
JOIN notePermission ON note.noteID = notePermission.noteID
JOIN noteUser ON notePermission.UserID = noteUser.userID
WHERE note.noteID = $1
`,
		noteID,
	)

	if err != nil {
		return note, err
	}

	for rows.Next() {
		var permission models.Permission
		rows.Scan(&permission.User.UserID, &permission.User.Username, &permission.User.Role, &permission.Permission)
		note.SharedUsers = append(note.SharedUsers, permission)
	}

	return note, nil
}

// columns are in the order: noteID, name, content, status, ownerID, dueDate, delegatedUserID
func (dbm *databasemanager) scanInNotes(rows *sql.Rows) ([]models.Note, error) {
	var notes = make([]models.Note, 0)
	for rows.Next() {
		note := models.Note{
			Owner:         &models.User{},
			DelegatedUser: &models.User{},
		}
		err := rows.Scan(&note.NoteID, &note.Name, &note.Content, &note.Status, &note.Owner.UserID, &note.DueDate, &note.DelegatedUser.UserID)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (dbm *databasemanager) GetNotesByOwner(userID int) ([]models.Note, error) {
	rows, err := dbm.db.Query(`
SELECT noteID, name, content, status, ownerID, dueDate, delegatedUserID
FROM note
WHERE ownerID = $1`, userID)

	if err != nil {
		return nil, err
	}

	return dbm.scanInNotes(rows)
}

func (dbm *databasemanager) GetNotesByDelegatedUser(userID int) ([]models.Note, error) {
	rows, err := dbm.db.Query(`
SELECT noteID, name, content, status, ownerID, dueDate, delegatedUserID
FROM note
WHERE delegatedUserID = $1`, userID)

	if err != nil {
		return nil, err
	}

	return dbm.scanInNotes(rows)
}

func (dbm *databasemanager) GetNotesSharedWith(userId int) ([]models.Note, error) {
	rows, err := dbm.db.Query(`
SELECT note.noteID, note.name, note.content, note.status, note.ownerID, note.dueDate, note.delegatedUserID
FROM notepermission
INNER JOIN note ON note.noteID = notepermission.noteID
WHERE notepermission.userID = $1
`, userId)

	if err != nil {
		return nil, err
	}

	return dbm.scanInNotes(rows)
}
