package common

import (
	"fmt"
)

type Node struct {
	NodeIp *string  `json:"node_ip"`
	NodeId *string `json:"node_id"`
}

func AddNode(n *Node) error {
	NodeList[n.NodeIp] = *n
}

func DeleteNode(n *Node) error {
	delete(NodeList, *n.NodeIp)
}

