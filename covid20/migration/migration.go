package migration

import (
	"summerclub20/covid20/database"
	"summerclub20/covid20/model"
)

//Migrate ... Database migration.
func Migrate() {
	database.DBCon.AutoMigrate(model.Daily{})
}
