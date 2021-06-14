package rpc

import (
	"net/url"

	"github.com/ethereum/go-ethereum/rpc"
)

type NetApi struct {
	rpcClient rpcClient
}

func NewNetApi(ip string, port string) (*NetApi, error) {
	rpcClient, err := rpc.Dial("http://" + ip + ":" + port)
	if err != nil {
		return nil, err
	}

	return &NetApi{
		rpcClient: rpcClient,
	}, nil
}

func (api *NetApi) NodeInfo() (*NodeInfo, error) {
	var enodeResponse rpcNetEnode
	err := api.rpcClient.Call(&enodeResponse, "net_enode")
	if err != nil {
		return nil, err
	}

	return parseEnode(string(enodeResponse))
}
