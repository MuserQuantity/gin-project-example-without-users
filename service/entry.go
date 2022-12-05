package service

import "github.com/MuserQuantity/gin-project-example-without-users/service/system"

type ServiceGroup struct {
	SystemServiceApp system.Group
}

var ServiceGroupApp = new(ServiceGroup)
