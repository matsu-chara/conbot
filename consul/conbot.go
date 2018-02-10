package consul

import (
	"sync"

	"github.com/hashicorp/consul/api"
)

// ConbotClient is client
type ConbotClient struct {
	consulClient *api.Client
}

// avoid too many call consul API via multi slack message
var consulMutex sync.Mutex

// Connect connect to consul
func Connect(address string) (*ConbotClient, error) {
	config := api.DefaultConfig()
	config.Address = address
	consulClient, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConbotClient{consulClient}, nil
}
