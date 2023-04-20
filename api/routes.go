package api

import (
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetApiRoutes(e *echo.Echo) {

	api := e.Group("/node")

	api.POST("/addnode", handlers.AddNode)

	api.POST("/deletenode", handlers.DeleteNode)

	g := e.Group("/*")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(common.Targets)))

}
