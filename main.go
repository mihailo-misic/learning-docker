package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	c "github.com/mihailo-misic/learning-docker/controllers"
	"github.com/mihailo-misic/learning-docker/database"
	m "github.com/mihailo-misic/learning-docker/models"
	method "github.com/bu/gin-method-override"
	"log"
	"github.com/gin-gonic/contrib/static"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	// Allow PATCH, PUT & DELETE method parsing
	r.Use(method.ProcessMethodOverride(r))
	// Load up static files
	r.Use(static.Serve("/static", static.LocalFile("./static", true)))
	// Use go html templates
	r.LoadHTMLGlob("views/*.gohtml")
	
	
	auth := r.Group("/auth")
	{
		auth.POST("/register", c.Register)
		auth.GET("/register", c.FormRegister)
		auth.POST("/login", c.Login)
		auth.GET("/login", c.FormLogin)
		auth.GET("/logout", c.Logout)
	}
	
	authorized := r.Group("/", c.IsAuth)
	{
		authorized.GET("/", c.GetProducts)
		users := authorized.Group("/users")
		{
			users.GET("/", c.GetUsers)
			users.POST("/", c.CreateUser)
			users.GET("/create", c.FormUser)
			users.GET("/show/:id", c.GetUser)
			users.GET("/edit/:id", c.FormUser)
			users.GET("/delete/:id", c.DeleteUser)
			users.PUT("/update/:id", c.UpdateUser)
		}
		
		products := authorized.Group("/products")
		{
			products.GET("/", c.GetProducts)
			products.POST("/", c.CreateProduct)
			products.GET("/create", c.FormProduct)
			products.GET("/show/:id", c.GetProduct)
			products.GET("/edit/:id", c.FormProduct)
			products.GET("/delete/:id", c.DeleteProduct)
			products.PUT("/update/:id", c.UpdateProduct)
		}
	}
	
	return r
}

func main() {
	// Setup database
	db := database.Init()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	
	db.AutoMigrate(&m.Product{})
	
	r := setupRouter()
	
	log.Fatal(r.Run(":8000"))
}
