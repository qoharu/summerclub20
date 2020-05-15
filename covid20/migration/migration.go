package migration

import (
	"covid20/database"
	"covid20/model"
)

//Migrate ... Database migration.
func Migrate() {
	database.DBCon.AutoMigrate(model.Daily{})
}
