package database

func (dbm *databasemanager) CreatePermission(noteID, userID int, permission string) error {
	_, err := dbm.db.Exec(
		`INSERT INTO notePermission (noteID, userID, permission) VALUES ($1, $2, $3)`,
		noteID,
		userID,
		permission,
	)

	return err
}

func (dbm *databasemanager) RemovePermission(noteId, userID int) error {
	_, err := dbm.db.Exec(
		`DELETE FROM notePermission
		WHERE noteID = $1 and userID = $2`,
		noteId,
		userID,
	)

	return err
}

func (dbm *databasemanager) UpdatePermission(noteID, userID int, permission string) error {
	_, err := dbm.db.Exec(
		`UPDATE notePermission
		SET permission = $3
		WHERE noteID = $1 and userID = $2`,
		noteID,
		userID,
		permission,
	)

	return err
}
