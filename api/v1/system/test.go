package system

import (
	"github.com/MuserQuantity/gin-project-example-without-users/dal/request"
	"github.com/MuserQuantity/gin-project-example-without-users/dal/response"
	"github.com/MuserQuantity/gin-project-example-without-users/service/system"
	"github.com/gin-gonic/gin"
)

type TestApi struct{}

func (api *TestApi) TestResponse(c *gin.Context) {
	var req request.Search
	err := c.ShouldBindJSON(&req)
	if err != nil {
		system.LogError(err)
		system.FailWithErrorArgs(c)
		return
	}
	system.SucceedWithData(c, response.TestResponse{Test: "haha"})
}
