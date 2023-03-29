package handlers

import (
	"github.com/labstack/echo"
)

func ProxyRequest(c echo.Context) error {
	common.lb(c.Response().Writer, c.Request())
	return nil
}
