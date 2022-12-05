package router

import "github.com/MuserQuantity/gin-project-example-without-users/router/system"

type Group struct {
	SystemRouterGroup system.RouterGroup
}

var GroupApp = &Group{}
