package migration

import (
	//"Auth/internal/database/gorm"
	//"Auth/internal/user/model"
	//"Auth/internal/database/gorm"
	"Auth/internal/database/gorm/config"
	"Auth/internal/user/model"
	"fmt"
)

func Migrate() error {
	db, err := config.GetDBcon()
	if err != nil {
		return err
	}
	db.AutoMigrate(model.User{})
	//db.AutoMigrate(model.Photo{})
	db.AutoMigrate(model.Address{})
	fmt.Print("Database successfully migrated")

	return err
}
