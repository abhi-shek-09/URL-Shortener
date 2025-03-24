package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"url-shortener/database"
	"url-shortener/handlers"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()
	
	r := gin.Default()

	// configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"}, // lets frontend access response headers
		AllowCredentials: true,                       // enables cookies or auth headers
		MaxAge:           12 * time.Hour,             // cache them for 12hrs
	}))


	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortcode", handlers.RetrieveURL)
	r.PUT("/:shortcode", handlers.UpdateURL)
	r.DELETE("/:shortcode", handlers.DeleteShortURL)
	r.GET("/:shortcode/stats", handlers.GenerateStats)

	r.Static("/static", "./static")

	r.Run(":8080")
}
