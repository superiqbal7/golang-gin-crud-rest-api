package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/superiqbal7/golang-gin-crud-rest-api/middlewares"
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

	server.Use(gin.Recovery(), middlewares.Logger())

	apiRoutes := server.Group("/api")
	routes.InitRoutes(apiRoutes)

	viewRoutes := server.Group("/view")
	routes.InitViewoutes(viewRoutes)

	server.Run(":8080")
}
