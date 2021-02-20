package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	// Login Endpoint: Authentication + Token creation
	router.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		fmt.Println(token)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})
}
