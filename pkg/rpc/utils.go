package rpc

import "net/url"

func parseEnode(enode string) (*NodeInfo, error) {

	url, err := url.Parse(enode)
	if err != nil {
		return nil, err
	}

	node := NodeInfo{
		Pubkey: url.User.String(),
		Port:   url.Port(),
	}

	return &node, nil
}
