package basic_learn

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGetEth(t *testing.T) {
	/*ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	ks.NewAccount("password")

	"0d518a4c445bbfad90c8382a051a91087d930253"
	"812b745ff0fdc32ecc024a6829a5908364d517f2"

	// LIMIT API KEY INFURA UNUTK GENERATE NEW URL UNUK ASES NETWORK LAIN, CARA NYA SAMA MASUKIN ADDRESS ETH AJA
	*/

	c, _ := ethclient.DialContext(ctx, ganacheURL)
	defer c.Close()

	a1 := common.HexToAddress("0x4EDFf68fE6f7074b71fA8342a04164249bFa6E39")
	a2 := common.HexToAddress("0xBb5D95F685b801A2C3d5fc3D6d43Ce4AA0dE848E")

	b1, err := c.BalanceAt(ctx, a1, nil)
	assert.Nil(t, err)
	b2, err := c.BalanceAt(ctx, a2, nil)
	assert.Nil(t, err)

	fmt.Println("balance addr 1 <-", b1)
	fmt.Println("balance addr 2 <-", b2)

}

func TestTransaction(t *testing.T) {

	// DATA WE USE
	var (
		privateKeySender   = "8b0fa74052aea23192a896e13ab81eda188f178d41f7a3dc7041096803f76148"
		publicKeyRecipient = common.HexToAddress("0xBb5D95F685b801A2C3d5fc3D6d43Ce4AA0dE848E")
		value              = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18)) // Mengirim 1 Ether (dalam Wei)
		gasLimit           = uint64(21000)
	)

	c, _ := ethclient.DialContext(ctx, ganacheURL)
	defer c.Close()

	privateKey, err := crypto.HexToECDSA(privateKeySender)
	assert.Nil(t, err)

	addressSender := crypto.PubkeyToAddress(privateKey.PublicKey)

	balanceSenderBefore, _ := c.BalanceAt(ctx, addressSender, nil)
	balanceRecipientBefore, _ := c.BalanceAt(ctx, publicKeyRecipient, nil)
	fmt.Println("Saldo pengirim sebelum transaksi:", balanceSenderBefore)
	fmt.Println("Saldo penerima sebelum transaksi:", balanceRecipientBefore)

	// AMBIL PUBLICT KEY SENDER
	getPubKeySender, err := crypto.HexToECDSA(privateKeySender)
	addrSender := crypto.PubkeyToAddress(getPubKeySender.PublicKey)
	assert.Nil(t, err)

	// EMNGAMBIL NONCE ATAU ANGKA (AUTO INC) SUDAH BERAPAKALI USER MELAKUKAN TX
	nonce, err := c.PendingNonceAt(ctx, addrSender)
	assert.Nil(t, err)

	// NGAMBIL CHAIN ID
	chainId, err := c.ChainID(context.Background())
	assert.Nil(t, err)

	//
	feeCap, _ := c.SuggestGasPrice(ctx)
	tipCap, _ := c.SuggestGasTipCap(ctx)

	// START TRANSACTION
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
		Gas:       gasLimit,
		To:        &publicKeyRecipient,
		Value:     value,
		Data:      nil,
	})

	signTx, err := types.SignTx(tx, types.NewLondonSigner(chainId), getPubKeySender)
	assert.Nil(t, err)

	err = c.SendTransaction(ctx, signTx)
	assert.Nil(t, err)

	balanceSender, _ := c.BalanceAt(ctx, addrSender, nil)
	balanceRecipient, _ := c.BalanceAt(ctx, publicKeyRecipient, nil)
	fmt.Println("ADDR SENDER AFTER <- ", balanceSender)
	fmt.Println("ADDR RECIPIENT AFTER <- ", balanceRecipient)
}
