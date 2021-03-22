package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Eth returns an ethereum client instance
type Eth struct {
	Client *ethclient.Client
}

// NewEthClient returns new ethereum client instance
func NewEthClient() (*Eth, error) {
	client, err := ethclient.Dial("https://kovan.infura.io/v3/6f02bb4df1f44aa39be13f3b9cddb18e")
	if err != nil {
		return nil, fmt.Errorf("Error '%v' while connecting to ethereum node", err)
	}
	return &Eth{Client: client}, nil
}
