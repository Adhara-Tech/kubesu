package rpc

type NodeInfo struct {
	Pubkey string
	Port   string
}

type rpcPeer struct {
	Version string     `json:"version"`
	Name    string     `json:"name"`
	Caps    []string   `json:"caps"`
	Network rpcNetwork `json:"network"`
	Port    string     `json:"port"`
	Id      string     `json:"id"`
}

type rpcNetwork struct {
	LocalAddress  string `json:"localAddress"`
	RemoteAddress string `json:"remoteAddress"`
}

type rpcNodeInfo struct {
	Enode      string         `json:"enode"`
	ListenAddr string         `json:"listenAddr"`
	Ip         string         `json:"ip"`
	Name       string         `json:"name"`
	Id         string         `json:"id"`
	Ports      map[string]int `ports:"id"`
	Protocols  interface{}    `json:"protocols"`
}
