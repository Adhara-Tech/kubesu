package types

import "fmt"

type Node struct {
	Name   string
	Ip     string
	Pubkey string
	Port   string
}

func (node *Node) Enode() string {
	return fmt.Sprintf("enode://%s@%s:%s", node.Pubkey, node.Ip, node.Port)
}
