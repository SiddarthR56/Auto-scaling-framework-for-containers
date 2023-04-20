package handlers

import (
	"net/http"
	"fmt"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"
	"github.com/labstack/echo"
)

func ProxyRequest(c echo.Context) error {
	fmt.Println("Request came")
	common.Lb(c.Response().Writer, c.Request())
	return c.JSON(http.StatusOK, "Message: Ok")
}
