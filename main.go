package main

import (
	"api"
	"net/http"
	"common"
	"github.com/labstack/echo"
)


func main() {

	e := echo.New()
	
	api.SetApiRoutes(e)

	e.Start(":8000")

	go common.healthCheck()

}