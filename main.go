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

	common.Node_pool.AddContainer("152.7.176.37", "30001", "30001")
	common.Node_pool.AddContainer("152.7.176.37", "30002", "30002")
	common.Node_pool.AddContainer("152.7.176.37", "30003", "30003")

	//go common.HealthCheck()

	e.Start(":8000")

}
