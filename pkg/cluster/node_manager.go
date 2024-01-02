package cluster

import (
	"fmt"
	"sync"

	"github.com/WuKongIM/WuKongIM/pkg/cluster/replica"
	"github.com/WuKongIM/WuKongIM/pkg/clusterevent/pb"
	"github.com/WuKongIM/WuKongIM/pkg/wkserver/proto"
)

type nodeManager struct {
	sync.RWMutex
	nodes map[uint64]*node
}

func newNodeManager() *nodeManager {
	return &nodeManager{
		nodes: make(map[uint64]*node),
	}
}

func (n *nodeManager) addNode(node *node) {
	n.Lock()
	defer n.Unlock()
	n.nodes[node.id] = node
}

func (n *nodeManager) removeNode(nodeID uint64) {
	n.Lock()
	defer n.Unlock()
	delete(n.nodes, nodeID)
}

func (n *nodeManager) getNode(nodeID uint64) *node {
	n.RLock()
	defer n.RUnlock()
	return n.nodes[nodeID]
}

func (n *nodeManager) getAllNode() []*node {
	n.RLock()
	defer n.RUnlock()
	nodes := make([]*node, 0, len(n.nodes))
	for _, node := range n.nodes {
		nodes = append(nodes, node)
	}
	return nodes
}

// 获取所有投票节点
func (n *nodeManager) getAllVoteNodes() []*node {
	n.RLock()
	defer n.RUnlock()
	nodes := make([]*node, 0, len(n.nodes))
	for _, node := range n.nodes {
		if node.allowVote {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (n *nodeManager) send(nodeID uint64, msg *proto.Message) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.send(msg)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) sendPing(nodeID uint64, req *PingRequest) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.sendPing(req)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) sendVote(nodeID uint64, req *VoteRequest) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.sendVote(req)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) sendVoteResp(nodeID uint64, req *VoteResponse) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.sendVoteResp(req)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) sendPong(nodeID uint64, req *PongResponse) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.sendPong(req)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) requestClusterConfig(nodeID uint64) (*pb.Cluster, error) {
	node := n.getNode(nodeID)
	if node != nil {
		return node.requestClusterConfig()
	}
	return nil, fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) requestSlotInfo(nodeID uint64, req *SlotInfoReportRequest) (*SlotInfoReportResponse, error) {
	node := n.getNode(nodeID)
	if node != nil {
		return node.requestSlotInfo(req)
	}
	return nil, fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) sendSyncNotify(nodeID uint64, req *replica.SyncNotify) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.sendSyncNotify(req)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) requestSyncLog(nodeID uint64, r *replica.SyncReq) (*replica.SyncRsp, error) {
	node := n.getNode(nodeID)
	if node != nil {
		return node.requestSyncLog(r)
	}
	return nil, fmt.Errorf("node[%d] not exist", nodeID)
}

func (n *nodeManager) requestAppendLog(nodeID uint64, req *AppendLogRequest) error {
	node := n.getNode(nodeID)
	if node != nil {
		return node.requestAppendLog(req)
	}
	return fmt.Errorf("node[%d] not exist", nodeID)

}
