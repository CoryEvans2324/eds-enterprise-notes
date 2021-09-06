package database

import (
	"github.com/CoryEvans2324/eds-enterprise-notes/models"
)

func (dbm *databasemanager) CreateUser(username, password string) (*models.User, error) {
	hash, _ := HashPassword(password)
	user := models.User{
		Username:     username,
		PasswordHash: hash,
	}
	result := dbm.db.Create(&user)

	return &user, result.Error
}

func (dbm *databasemanager) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := dbm.db.First(&user, userID)
	return &user, result.Error
}

func (dbm *databasemanager) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := dbm.db.Where("Username = ?", username).First(&user)
	return &user, result.Error
}

func (dbm *databasemanager) SearchForUsername(username string) ([]string, error) {
	var usernames = make([]string, 0)
	result := dbm.db.Model(&models.User{}).Where("LOWER(Username) LIKE LOWER('%' || ? || '%')", username).Pluck("Username", &usernames)

	return usernames, result.Error
}
