package common

import "sync/atomic"

type Node struct {
	NodeIp *string `json:"node_ip"`
	NodeId *string `json:"node_id"`
}

type NodeList struct {
	nodes  []*Node
	curInd uint64
}

func AddNodeInt(n Node) error {
	Node_Map[*n.NodeIp] = n
	Node_List.nodes = append(Node_List.nodes, &n)
	return nil
}

func (s *NodeList) GetNextPeer() *Node {
	if len(s.nodes) != 0 {
		next := int(atomic.AddUint64(&s.curInd, uint64(1)) % uint64(len(s.nodes)))

		atomic.StoreUint64(&s.curInd, uint64(next))

		return s.nodes[next]
	}
	return nil
}

func DeleteNodeInt(n Node) error {
	delete(Node_Map, *n.NodeIp)
	for i, b := range Node_List.nodes {
		if *b.NodeIp == *n.NodeIp {
			Node_List.nodes = append(Node_List.nodes[:i], Node_List.nodes[i+1:]...)
		}
	}
	return nil
}
