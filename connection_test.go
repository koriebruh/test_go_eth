package go_etherium_learn

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.Background()
var infraURL = "https://mainnet.infura.io/v3/7108f6b019944d2082df7b667e6b1f4a"
var ganacheURL = "http://localhost:8545"

func TestConnectionOnline(t *testing.T) {
	client, err := ethclient.DialContext(ctx, infraURL)
	assert.Nil(t, err)
	defer client.Close()

	block, err := client.BlockByNumber(ctx, nil)
	assert.Nil(t, err)

	fmt.Printf("Block Number: %d\nTransactions: %+v\n", block.Number(), block.Transactions())
	assert.NotNil(t, block)
}

func TestConnectionLocalhost(t *testing.T) {
	client, err := ethclient.DialContext(ctx, ganacheURL)
	assert.Nil(t, err)
	defer client.Close()

	block, err := client.BlockByNumber(ctx, nil)
	assert.Nil(t, err)

	fmt.Printf("Block Number: %d\nTransactions: %+v\n", block.Number(), block.Transactions())
	assert.NotNil(t, block)
}
