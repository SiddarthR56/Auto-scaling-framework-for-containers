package handlers

import (
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"
	"github.com/labstack/echo"
)

func ProxyRequest(c echo.Context) error {
	common.Lb(c.Response().Writer, c.Request())
	return nil
}
