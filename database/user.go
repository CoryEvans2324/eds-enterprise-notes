package database

import "github.com/CoryEvans2324/eds-enterprise-notes/models"

func (dbm *databasemanager) CreateUser(username, password string) (int, error) {
	hash, _ := HashPassword(password)
	row := dbm.db.QueryRow(`INSERT INTO noteUser (username, passwordHash) VALUES ($1, $2) RETURNING userID`, username, hash)

	var userID int
	err := row.Scan(&userID)

	return userID, err
}

func (dbm *databasemanager) GetUserByID(userID int) (*models.User, error) {
	row := dbm.db.QueryRow(`SELECT username, userRole FROM noteUser WHERE userID = $1`, userID)

	user := models.User{
		UserID: userID,
	}

	err := row.Scan(&user.Username, &user.Role)

	return &user, err
}

func (dbm *databasemanager) GetUserByUsername(username string) (*models.User, error) {
	row := dbm.db.QueryRow(`SELECT userID, userRole FROM noteUser WHERE username = $1`, username)

	user := models.User{
		Username: username,
	}

	err := row.Scan(&user.UserID, &user.Role)

	return &user, err
}

func (dbm *databasemanager) GetPasswordHash(username string) (string, error) {
	row := dbm.db.QueryRow(`SELECT passwordHash FROM noteUser WHERE username = $1`, username)

	var hash string
	err := row.Scan(&hash)

	return hash, err
}
