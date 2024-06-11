package main

import (
	"os"

	"github.com/ImArnav19/mongo/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/create", routes.Create)
	router.GET("/read", routes.Read)
	router.PUT("/update/:id", routes.Update)
	router.DELETE("/del/:id", routes.Delete)
	router.Run(":" + port)
}
