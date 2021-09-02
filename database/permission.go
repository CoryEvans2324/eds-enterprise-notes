package database

import "github.com/CoryEvans2324/eds-enterprise-notes/models"

func (dbm *databasemanager) CreatePermission(permission models.Permission) error {
	result := dbm.db.Create(&permission)
	return result.Error
}

func (dbm *databasemanager) RemovePermission(permission models.Permission) error {
	result := dbm.db.Delete(&permission)
	return result.Error
}

func (dbm *databasemanager) UpdatePermission(permission models.Permission) error {
	result := dbm.db.Save(&permission)
	return result.Error
}
