package routes

import "github.com/gin-gonic/gin"

func SetViewRoutes(router *gin.RouterGroup) {

	router.GET("/videos", videoController.ShowAll)
}
