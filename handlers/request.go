package handlers

import (
	"net/http"

	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/contracts"
	"github.com/labstack/echo"
)

func ProxyRequest(c echo.Context) error {
	common.Lb(c.Response().Writer, c.Request())
	return nil
}

func ContainerAdd(c echo.Context) error {

	params := new(contracts.ContainerModifyRequest)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	Application := params.AppName
	ContainerNumber := params.ContainerNum

	for i := 0; i < *ContainerNumber; i++ {
		node := common.Node_List.GetNextPeer()
		CreateContainer(*Application, *node.NodeIp, common.WNODE_PORT)
	}

	message := "Container was Added"
	response := contracts.BaseResponse{}
	response.Message = &message
	return c.JSON(http.StatusOK, response)

}

func ContainerDelete(c echo.Context) error {
	params := new(contracts.ContainerModifyRequest)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	Application := params.AppName
	ContainerNumber := params.ContainerNum

	for i := 0; i < *ContainerNumber; i++ {
		node := common.Node_List.GetNextPeer()
		//common.Node_List.DeleteContainer(*node.NodeIp, *node)
		DeleteContainer(*Application, *node.NodeIp, common.WNODE_PORT)
	}

	message := "Container was Deleted"
	response := contracts.BaseResponse{}
	response.Message = &message
	return c.JSON(http.StatusOK, response)

}
