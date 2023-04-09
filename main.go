package main

import (
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/api"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	api.SetApiRoutes(e)

	common.Node_pool.AddContainer("127.0.0.1", "2000")

	e.Start(":8000")

	//go common.HealthCheck()

}
