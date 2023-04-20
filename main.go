package main

import (
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/api"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"

	"github.com/labstack/echo"
)

// Main function
func main() {

	e := echo.New()

	api.SetApiRoutes(e)

	common.AddContainer("152.7.176.37", "30001")
	common.AddContainer("152.7.176.37", "30002")
	common.AddContainer("152.7.176.37", "30003")

	e.Start(":8000")

	//go common.HealthCheck()

}
