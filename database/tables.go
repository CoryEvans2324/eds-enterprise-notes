package database

func (dbm *databasemanager) CreateUserTable() error {
	_, err := dbm.db.Exec(`
CREATE TABLE noteUser (
	userID SERIAL PRIMARY KEY,
	username varchar UNIQUE,
	passwordHash varchar,
	userRole varchar
)
`)
	return err
}

func (dbm *databasemanager) DropUserTable() error {
	_, err := dbm.db.Exec(`DROP TABLE IF EXISTS noteUser`)
	return err
}
