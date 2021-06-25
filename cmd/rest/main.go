package main

import (
	"Auth/cmd/rest/router"
	"Auth/internal/database/gorm/config"
	"Auth/internal/migration"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("C:\\Users\\Tilefamily\\Documents\\AuthHexa\\internal\\env\\.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	siteTitle := os.Getenv("SITE_TITLE")
	dbHost := os.Getenv("DB_HOST")

	fmt.Printf("godotenv : %s = %s \n", "Site Title", siteTitle)
	fmt.Printf("godotenv : %s = %s \n", "DB Host", dbHost)
	////initialize the database
	config.GetDBcon()

	migration.Migrate() ///add this line to main.go to initialize the migration
	db, err := config.GetDBcon()
	if err != nil {
		print(err)
	}
	dbc, err := db.DB()
	///finally close the connection when you are done
	defer dbc.Close()
	log.Fatalln(http.ListenAndServe(":8080", router.NewRouter()))

}
