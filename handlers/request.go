package handlers

import (
	"fmt"
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

	fmt.Println("Application Name: ", *Application)

	for i := 0; i < *ContainerNumber; i++ {
		node := common.Node_List.GetNextPeer()
		NodeCreateContainer(*Application, *node.NodeIp, common.WNODE_PORT)
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

	// Application := params.AppName
	// pool:= common.AppNameMap[*Application]
	ContainerNumber := params.ContainerNum

	for i := 0; i < *ContainerNumber; i++ {
		node := common.Node_pool.GetNextPeer()
		common.Node_pool.DeleteSpecificContainer(node.ContainerID)

		go NodeDeleteContainer(node.ContainerID, common.ContainerList[node.ContainerID], common.WNODE_PORT)
	}

	message := "Container was Deleted"
	response := contracts.BaseResponse{}
	response.Message = &message
	return c.JSON(http.StatusOK, response)

}

func ContainerRestart(c echo.Context) error {

	params := new(contracts.ContainerRestartRequest)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	containerId := params.ContainerId

	common.Node_pool.DeleteSpecificContainer(*containerId)

	node := common.Node_List.GetNextPeer()

	NodeCreateContainer("rubis", *node.NodeIp, common.WNODE_PORT)

	go NodeDeleteContainer(*containerId, common.ContainerList[*containerId], common.WNODE_PORT)

	return nil
}

func AddApplication(c echo.Context) error {

	params := new(contracts.AppAddRequest)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ApplicationName := params.AppName
	ApplicationImage := params.AppImage

	common.AppImageMap[*ApplicationName] = *ApplicationImage

	fmt.Println("Adding Application: ", *ApplicationName)

	node := common.Node_List.GetNextPeer()
	NodeCreateContainer(*ApplicationImage, *node.NodeIp, common.WNODE_PORT)

	message := "Application was Added"
	response := contracts.BaseResponse{}
	response.Message = &message
	return c.JSON(http.StatusOK, response)

}
