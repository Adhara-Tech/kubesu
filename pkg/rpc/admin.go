package rpc

import (
	"net/url"

	"github.com/ethereum/go-ethereum/rpc"
)

type AdminApi struct {
	rpcClient *rpc.Client
}

func NewAdminApi(ip string, port string) (*AdminApi, error) {
	rpcClient, err := rpc.Dial("http://" + ip + ":" + port)
	if err != nil {
		return nil, err
	}

	return &AdminApi{
		rpcClient: rpcClient,
	}, nil
}

func (api *AdminApi) NodeInfo() (*NodeInfo, error) {
	var nodeInfo rpcNodeInfo
	err := api.rpcClient.Call(&nodeInfo, "admin_nodeInfo")
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(nodeInfo.Enode)
	if err != nil {
		return nil, err
	}

	node := NodeInfo{
		Pubkey: url.User.String(),
		Port:   url.Port(),
	}

	return &node, nil
}

func (api *AdminApi) Peers() ([]NodeInfo, error) {
	var rpcPeers []rpcPeer
	err := api.rpcClient.Call(&rpcPeers, "admin_peers")
	if err != nil {
		return nil, err
	}

	var nodes []NodeInfo
	for _, peer := range rpcPeers {
		node := NodeInfo{
			Pubkey: peer.Id[2:],
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (api *AdminApi) AddPeer(enode string) (bool, error) {
	var success bool
	err := api.rpcClient.Call(&success, "admin_addPeer", enode)
	if err != nil {
		return false, err
	}

	return success, nil
}
