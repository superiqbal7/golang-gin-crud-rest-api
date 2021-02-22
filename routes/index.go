package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/superiqbal7/golang-gin-crud-rest-api/controller"
	"github.com/superiqbal7/golang-gin-crud-rest-api/middlewares"
	"github.com/superiqbal7/golang-gin-crud-rest-api/repository"
	"github.com/superiqbal7/golang-gin-crud-rest-api/service"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func InitRoutes(g *gin.RouterGroup) {

	g.Use(middlewares.AuthorizeJWT())
	//g.Use(middlewares.BasicAuth(), gindump.Dump())
	SetVideoRoutes(g)
	//SetAuthRoutes(g)
}

func InitViewoutes(g *gin.RouterGroup) {
	SetViewRoutes(g)
}

func InitAuthRoutes(g *gin.RouterGroup) {
	fmt.Println(g)
	AuthRoutes(g)
}
