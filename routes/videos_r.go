package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetVideoRoutes(router *gin.RouterGroup) {
	// @api {get} /api/videos

	router.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	router.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		}
	})
}
