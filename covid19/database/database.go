package database

import "github.com/jinzhu/gorm"

var (
	// DBCon ... Database Connection Instance
	DBCon *gorm.DB
)

// InitDB ... function to initialize database
func InitDB() {
	var err error

	//connect to postgres database
	DBCon, err = gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=covid19 password=rahasia123 sslmode=disable")

	if err != nil {
		panic(err)
	}

}
