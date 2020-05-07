package migration

import (
	"summerclub20/covid19/database"
	"summerclub20/covid19/model"
)

//Migrate ... Database migration.
func Migrate() {
	database.DBCon.AutoMigrate(model.Daily{})
}
