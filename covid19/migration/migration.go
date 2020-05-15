package migration

import (
	"covid19/database"
	"covid19/model"
)

//Migrate ... Database migration.
func Migrate() {
	database.DBCon.AutoMigrate(model.Daily{})
}
