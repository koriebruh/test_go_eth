package basic_learn

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

func TestInteract(t *testing.T) {
	client, _ := ethclient.DialContext(ctx, infraURL)
	defer client.Close()

	//RANDOM Address
	address := common.HexToAddress("0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5")
	balanceWei, _ := client.BalanceAt(ctx, address, nil)
	weiToEther := new(big.Float).SetFloat64(1e18)
	balanceEther := new(big.Float).Quo(new(big.Float).SetInt(balanceWei), weiToEther)
	fmt.Printf("Balance: %s Wei (%s ETH)\n", balanceWei.String(), balanceEther.Text('f', 18))

	//

}
