package api

import (
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/handlers"
	"github.com/labstack/echo"
)

func SetApiRoutes(e *echo.Echo) {

	api := e.Group("/admin")

	api.POST("/addnode", handlers.AddNode)

	api.POST("/addapplication", handlers.AddApplication)

	api.POST("/deletenode", handlers.DeleteNode)

	api.POST("/containerrestart", handlers.ContainerRestart)

	api.POST("/containeradd", handlers.ContainerAdd)

	api.POST("/containerdelete", handlers.ContainerDelete)

	e.GET("/*", handlers.ProxyRequest)
	e.POST("/*", handlers.ProxyRequest)

}
