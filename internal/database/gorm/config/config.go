package config

import (
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func GetDBcon() (*gorm.DB, error) {
	var err error

	dsn := "host=localhost user=postgres password=dbpass dbname=Usermanagment port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dsn == "true" {
		print("Successfully connected to database")
	}

	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}
