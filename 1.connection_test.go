package go_etherium_learn

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	fmt.Printf("Block Number: %d\n", block.Number())
	fmt.Printf("Block Number: %d\n", block.Number())
	fmt.Printf("Timestamp: %d\n", block.Time())
	fmt.Printf("Number of Transactions: %d\n\n", len(block.Transactions()))
	for _, tx := range block.Transactions() {
		fmt.Printf("Tx Hash: %s\n", tx.Hash().Hex())
	}
	assert.NotNil(t, block)
}
