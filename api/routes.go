package api

import (

	"fmt"
	"handlers"
	"github.com/labstack/echo/v4"

)

func SetApiRoutes (e *echo.Echo) {

	api := e.Group("/api")

	api.POST("/addnode", handlers.AddNode)

	api.POST("/deletenode", handlers.DeleteNode)

	api.POST("/frequest", handlers.ProxyRequest)

}