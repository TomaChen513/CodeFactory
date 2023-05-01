package router

import (
	"thirdPartLogin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine{
	router:=gin.Default()

	router.GET("/",controller.Login)
	router.GET("/github_login", controller.HandlerGithubLogin)
	router.GET("/callbackGithub", controller.GetGithubToken)

	return router
}