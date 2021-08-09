package database

func (dbm *databasemanager) CreateUserTable() error {
	_, err := dbm.db.Exec(`
CREATE TABLE noteUser (
	userID SERIAL PRIMARY KEY,
	username varchar NOT NULL UNIQUE,
	passwordHash varchar NOT NULL,
	userRole varchar DEFAULT 'user'
)
`)
	return err
}

func (dbm *databasemanager) DropUserTable() error {
	_, err := dbm.db.Exec(`DROP TABLE IF EXISTS noteUser`)
	return err
}
