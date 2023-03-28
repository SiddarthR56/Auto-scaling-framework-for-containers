package handlers

import (
	"fmt"
	"common"
	"contracts"
	"github.com/labstack/echo"
)

func AddNode(c echo.Context) error {
	params := new(contracts.NodeAddRequest)

	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	n := common.Node{}
	n.NodeIp = params.NodeIp
	n.NodeId = params.NodeId

	common.AddNode(n)

	message := "Node was Added"
	response := contracts.BaseResponse{}
	response.Message = &message
	return ctx.JSON(http.StatusOK, response)
}

func Deletenode(c echo.Context) error {
	params := new(contracts.NodeDeleteRequest)	

	if err := ctx.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	n := common.Node{}
	n.NodeIp = params.NodeIp
	n.NodeId = params.NodeId

	common.DeleteNode(n)

	message := "Node was Deleted"
	response := contracts.BaseResponse{}
	response.Message = &message
	return ctx.JSON(http.StatusOK, response)

}