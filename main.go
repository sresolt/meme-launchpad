package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/janction/meme-launchpad/database"
	"github.com/janction/meme-launchpad/routes"
)

func main() {
	dsn := "host=localhost user=postgres password=tu_contraseña_segura dbname=memelaunchpad port=5432 sslmode=disable TimeZone=UTC"

	if err := database.Connect(dsn); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	if err := database.Migrate(database.DB); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	r := gin.Default()
	routes.RegisterTokenRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
