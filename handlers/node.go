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

	common.Node_pool.AddContainer(nodeip, port)

	return nil

}
