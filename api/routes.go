package api

import (
	"net/url"

	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetApiRoutes(e *echo.Echo) {

	url1, _ := url.Parse("http://localhost:8081")

	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}

	api := e.Group("/api")

	api.POST("/addnode", handlers.AddNode)

	api.POST("/deletenode", handlers.DeleteNode)

	api.POST("/frequest", handlers.ProxyRequest)

	g := e.Group("/blog")
	g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

}
