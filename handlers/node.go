package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/common"
	"github.com/SiddarthR56/Auto-scaling-framework-for-containers/contracts"

	"github.com/labstack/echo"
)

func AddNode(c echo.Context) error {
	params := new(contracts.NodeAddRequest)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	n := common.Node{}
	n.NodeIp = params.NodeIp
	n.NodeId = params.NodeId

	common.AddNode(n)

	message := "Node was Added"
	response := contracts.BaseResponse{}
	response.Message = &message
	return c.JSON(http.StatusOK, response)
}

func DeleteNode(c echo.Context) error {
	params := new(contracts.NodeDeleteRequest)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	n := common.Node{}
	n.NodeIp = params.NodeIp
	n.NodeId = params.NodeId

	common.DeleteNode(n)

	message := "Node was Deleted"
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

	request := map[string]string{
		"container_id": *containerId,
	}

	req, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = common.MakePostRequest(common.ContainerList[*containerId], req)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

func CreateContainer(funcName string, numContainers string, nodeip string, nodeport string) error {

	url := fmt.Sprintf("http://%s:%s/api/createcontainer", nodeip, nodeport)

	request := map[string]string{
		"function_name": funcName,
		"numContainers": numContainers,
	}

	req, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := common.MakePostRequest(url, req)
	if err != nil {
		fmt.Println(err.Error())
	}

	port := result["port"].(string)
	containerId := result["container_id"].(string)

	common.ContainerList[containerId] = fmt.Sprintf("http://%s:%s/api/restartcontainer", nodeip, nodeport)

	common.Node_pool.AddContainer(nodeip, port)

	return nil

}
