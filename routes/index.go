package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/superiqbal7/golang-gin-crud-rest-api/controller"
	"github.com/superiqbal7/golang-gin-crud-rest-api/middlewares"
	"github.com/superiqbal7/golang-gin-crud-rest-api/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func InitRoutes(g *gin.RouterGroup) {
	g.Use(middlewares.BasicAuth(), gindump.Dump())
	SetVideoRoutes(g)
	//SetAuthRoutes(g)
}

func InitViewoutes(g *gin.RouterGroup) {
	SetViewRoutes(g)
}
