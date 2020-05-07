package migration

import (
	"qoharu/summerclub20/covid19/database"
	"qoharu/summerclub20/covid19/model"
)

//Migrate ... Database migration.
func Migrate() {
	database.DBCon.AutoMigrate(model.Daily{})
}
