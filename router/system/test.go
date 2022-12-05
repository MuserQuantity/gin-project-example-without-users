package system

import (
	v1 "github.com/MuserQuantity/gin-project-example-without-users/api/v1"
	"github.com/gin-gonic/gin"
)

type TestRouter struct{}

func (t *TestRouter) InitTestRouter(Router *gin.RouterGroup) {
	testRouter := Router.Group("test")
	testApi := v1.ApiGroupApp.SystemApiGroup.TestApi
	{
		testRouter.POST("test", testApi.TestResponse)
	}
}
