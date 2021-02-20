package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/superiqbal7/golang-gin-crud-rest-api/routes"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), gin.Logger())

	authRoutes := server.Group("/auth")
	routes.InitAuthRoutes(authRoutes)

	apiRoutes := server.Group("/api")
	routes.InitRoutes(apiRoutes)

	viewRoutes := server.Group("/view")
	routes.InitViewoutes(viewRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
