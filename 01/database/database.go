package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	m "github.com/mihailo-misic/learning-docker/01/models"
	c "github.com/mihailo-misic/learning-docker/01/controllers"
	"log"
)

func Init() (db *gorm.DB) {
	// Open a db connection
	var err error
	db, err = gorm.Open("postgres", `
		host=postgres
		port=5433
		user=mygo
		dbname=crm
		password=secret
		sslmode=disable
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&m.Product{})

	// Connect the controllers to the Database
	c.ConnectControllers(db)

	return
}
