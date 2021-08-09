package database

func (dbm *databasemanager) CreateUserTable() error {
	_, err := dbm.db.Exec(`
CREATE TABLE noteUser (
	userID SERIAL PRIMARY KEY,
	username VARCHAR NOT NULL UNIQUE,
	passwordHash VARCHAR NOT NULL,
	userRole VARCHAR DEFAULT 'user'
)
`)
	return err
}

func (dbm *databasemanager) DropUserTable() error {
	_, err := dbm.db.Exec(`DROP TABLE IF EXISTS noteUser`)
	return err
}

func (dbm *databasemanager) CreateNoteTable() error {
	_, err := dbm.db.Exec(`
CREATE TABLE note (
	noteID SERIAL PRIMARY KEY,
	name VARCHAR,
	content VARCHAR,
	status VARCHAR,
	ownerID INTEGER,
	dueDate TIMESTAMP,
	delegatedUserID INTEGER DEFAULT -1
)
`)
	return err
}

func (dbm *databasemanager) DropNoteTable() error {
	_, err := dbm.db.Exec(`DROP TABLE IF EXISTS note`)
	return err
}

func (dbm *databasemanager) CreatePermissionTable() error {
	_, err := dbm.db.Exec(`
CREATE TABLE notePermission (
	noteID INTEGER,
	userID INTEGER,
	permission VARCHAR
)
`)
	return err
}

func (dbm *databasemanager) DropPermissionTable() error {
	_, err := dbm.db.Exec(`DROP TABLE IF EXISTS notePermission`)
	return err
}
