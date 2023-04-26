package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	common.AddNodeInt(n)

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

	common.DeleteNodeInt(n)

	message := "Node was Deleted"
	response := contracts.BaseResponse{}
	response.Message = &message
	return c.JSON(http.StatusOK, response)

}

func NodeCreateContainer(funcName string, nodeip string, nodeport string) error {

	url := fmt.Sprintf("http://%s:%s/api/createcontainer", nodeip, nodeport)

	request := map[string]string{
		"function_name": funcName,
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

	common.ContainerList[containerId] = nodeip

	common.Node_pool.AddContainer(nodeip, port, containerId)

	return nil

}

func NodeDeleteContainer(container_id string, nodeip string, nodeport string) error {

	time.Sleep(30 * time.Second)

	url := fmt.Sprintf("http://%s:%s/api/deletecontainer", nodeip, nodeport)

	request := map[string]string{
		"container_id": container_id,
	}

	req, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = common.MakePostRequest(url, req)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil

}
