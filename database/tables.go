package database

import "github.com/CoryEvans2324/eds-enterprise-notes/models"

func (dbm *databasemanager) AutoMigrate() {
	dbm.db.AutoMigrate(&models.User{})
	dbm.db.AutoMigrate(&models.Note{})
	dbm.db.AutoMigrate(&models.Permission{})
}

func (dbm *databasemanager) DropTables() {
	dbm.db.Migrator().DropTable(&models.User{})
	dbm.db.Migrator().DropTable(&models.Note{})
	dbm.db.Migrator().DropTable(&models.Permission{})
}
