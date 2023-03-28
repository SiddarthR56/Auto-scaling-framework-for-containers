package handlers

import (
	"net/http"
	"common"
)

func ProxyRequest(c echo.Context) error {
	common.lb(c.Response().Writer, c.Request())
}