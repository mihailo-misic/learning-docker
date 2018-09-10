package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/mihailo-misic/learning-docker/controllers"
	. "github.com/mihailo-misic/learning-docker/models"
	"log"
)

func Init() (db *gorm.DB) {
	// Open a db connection
	var err error
	db, err = gorm.Open("postgres", `
		host=postgres
		port=5432
		user=mygo
		dbname=crm
		password=secret
		sslmode=disable
`)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{}, &User{})

	// Connect the controllers to the Database
	ConnectControllers(db)

	return
}
