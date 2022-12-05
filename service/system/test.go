package system

import (
	"github.com/MuserQuantity/gin-project-example-without-users/dal/response"
)

type TestService struct{}

var TestServiceApp = &TestService{}

func (t *TestService) TestFunc(arg string) (resp response.TestResponse, err error) {
	return response.TestResponse{Test: "haha"}, nil
}
