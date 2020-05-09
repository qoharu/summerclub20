package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	// DBCon ... Database Connection Instance
	DBCon *gorm.DB
)

// InitDB ... function to initialize database
func InitDB() {
	var err error

	var dbUser string = "postgres"
	var dbPass string = "rahasia123"
	var dbName string = "covid19"

	dbString := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPass)

	//connect to postgres database
	DBCon, err = gorm.Open("postgres", dbString)

	if err != nil {
		panic(err)
	}

}
