package system

import "github.com/MuserQuantity/gin-project-example-without-users/service"

type ApiGroup struct {
	TestApi
}

var (
	testService = service.ServiceGroupApp.SystemServiceApp.TestService
)
