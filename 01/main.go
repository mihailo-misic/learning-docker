package main

import (
	"log"
	"github.com/gin-gonic/gin"
	c "github.com/mihailo-misic/learning-docker/01/controllers"
	"github.com/mihailo-misic/learning-docker/01/database"
)

var r = gin.Default()

func main() {
	// Setup database
	db := database.Init()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	// Get all
	r.GET("/", c.GetProducts)
	// Get one
	// Create
	// Update
	// Delete

	log.Fatal(r.Run(":8000"))
}
