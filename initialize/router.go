package initialize

import (
	"github.com/gin-gonic/gin"

	"github.com/MuserQuantity/gin-project-example-without-users/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouterGroup := router.GroupApp.SystemRouterGroup
	ApiGroup := Router.Group("api")
	{
		systemRouterGroup.TestRouter.InitTestRouter(ApiGroup)
	}
	return Router
}
