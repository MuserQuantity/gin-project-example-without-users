package v1

import "github.com/MuserQuantity/gin-project-example-without-users/api/v1/system"

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = &ApiGroup{}
