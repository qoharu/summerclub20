package migration

import (
	"summerclub20/skeleton/database"
	"summerclub20/skeleton/model"
)

//Migrate ... Database migration.
func Migrate() {
	database.DBCon.AutoMigrate(model.Daily{})
}
