package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	lendingContract "koriebruh/learn/gen"
	"log"
	"math/big"
	"os"
	"testing"
)

var (
	ctx      = context.Background()
	infraURL = "https://sepolia.infura.io/v3/7108f6b019944d2082df7b667e6b1f4a"
	passOwn  = "jamal" // owner contract pass
	passAdd1 = "password"
	cAddr    = common.HexToAddress("0x90Bf7bEd8846C9449A9a66CaE8363C614f5d9cd1")
)

// TEST ISSUE LOAN FROR ADD Xxxx
func TestOwnerAddAddr(t *testing.T) {
	//OWNER ADD ADDR TO ALLOW LENDING
	path := "E:/Development/GO DEV/go-etherium-learn/wallet/UTC--2025-01-14T11-02-41.636529300Z--0d518a4c445bbfad90c8382a051a91087d930253"
	fmt.Printf("Reading file from: %s\n", path)

	// Baca file
	file, err := os.ReadFile(path)
	ifErrLog(err, "can't read file")
	decryptKey, err := keystore.DecryptKey(file, passOwn)
	ifErrLog(err, "can't decrypt")

	// EXTRACT FROM FILE
	pvKey := decryptKey.PrivateKey
	//pvKeyBytes := crypto.FromECDSA(pvKey)
	//pubKey := decryptKey.PrivateKey.PublicKey
	//pubKeyBytes := crypto.FromECDSAPub(&pubKey)
	//addr := crypto.PubkeyToAddress(pubKey).Hex()

	c, err := ethclient.DialContext(ctx, infraURL)
	defer c.Close()
	ifErrLog(err, "fail to connect infra")

	// CHEcK CA
	code, err := c.CodeAt(ctx, cAddr, nil)
	ifErrLog(err, "fetching contract code")
	if len(code) == 0 {
		log.Fatalf("No contract code found at address: %s\n", cAddr.Hex())
	}
	fmt.Println("Contract code found at address:", cAddr.Hex())

	networkID, err := c.NetworkID(ctx)
	ifErrLog(err, "can't fetch network ID")
	tipCap, _ := c.SuggestGasTipCap(ctx)
	baseFee := big.NewInt(20000000)
	gasFeeCap := new(big.Int).Add(baseFee, tipCap)

	lc, err := lendingContract.NewLendingContract(cAddr, c)
	ifErrLog(err, "can't find contract")

	tx, err := bind.NewKeyedTransactorWithChainID(pvKey, networkID)
	ifErrLog(err, "err lag")

	tx.GasTipCap = tipCap
	tx.GasFeeCap = gasFeeCap

	borrowerAddr := common.HexToAddress("0x812b745ff0fdc32ecc024a6829a5908364d517f2")
	issueLoan, err := lc.IssueLoan(tx, borrowerAddr, big.NewInt(90000))
	ifErrLog(err, "error issue loan")

	fmt.Println("yey success issue loan for <- ", borrowerAddr)
	fmt.Println("success hash tx <- ", issueLoan.Hash())

}

func TestPublicMethod(t *testing.T) {

	borrowerAddr := common.HexToAddress("0x812b745ff0fdc32ecc024a6829a5908364d517f2")

	// connect infra
	c, err := ethclient.DialContext(ctx, infraURL)
	assert.Nil(t, err)

	// call contract
	contract, err := lendingContract.NewLendingContract(cAddr, c)
	assert.Nil(t, err)

	opts := &bind.CallOpts{
		Pending: true,                    // Gunakan nilai true untuk memanggil status transaksi terkini
		From:    common.HexToAddress(""), // Opsional, bisa nil jika tidak spesifik
	}

	loanAmount, interestRate, dueDate, paidStatus, err := contract.CheckLoanStatus(opts, borrowerAddr)
	assert.Nil(t, err)
	t.Logf("\nLOAN AMOUNT := %v\nLOAN INTEREST RATE := %v\nLOAN DUE-DATE := %d\nPAID STATUS := %v\n", loanAmount, interestRate, dueDate, paidStatus)

	overdueLoan, err := contract.CheckOverdueLoan(opts, borrowerAddr)
	assert.Nil(t, err)
	t.Logf("APKAH SUDAH LEWAT TANGGAL <- %v", overdueLoan)

}

func TestPayLoan(t *testing.T) {
	path := "wallet/UTC--2025-01-14T23-38-05.689713800Z--812b745ff0fdc32ecc024a6829a5908364d517f2"
	file, err := os.ReadFile(path)
	assert.Nil(t, err)

	decryptKey, err := keystore.DecryptKey(file, passAdd1)
	assert.Nil(t, err)
	pvKey := decryptKey.PrivateKey
	pubKey := decryptKey.PrivateKey.PublicKey
	//pubKeyBytes := crypto.FromECDSAPub(&pubKey)
	addr := crypto.PubkeyToAddress(pubKey)

	// CONNECT ETH CLIENT
	c, err := ethclient.DialContext(ctx, infraURL)
	assert.Nil(t, err)

	// PARAM NEED
	networkID, err := c.NetworkID(ctx)
	ifErrLog(err, "can't fetch network ID")
	nonceAt, err := c.PendingNonceAt(ctx, addr)
	assert.Nil(t, err)

	//tipCap, _ := c.SuggestGasTipCap(ctx)
	tipCap := big.NewInt(50000000000)
	baseFee := big.NewInt(2000000000) // 20 Gwei
	gasFeeCap := new(big.Int).Add(baseFee, tipCap)

	// CEK BALANCE BORROWER
	balance, err := c.BalanceAt(ctx, addr, nil)
	assert.Nil(t, err)
	t.Logf("Current balance: %s wei (address: %s)", balance.String(), addr.Hex())

	// MAKE Tx FOR PAYABLE FUNC
	auth, err := bind.NewKeyedTransactorWithChainID(pvKey, networkID)
	assert.Nil(t, err)
	auth.GasTipCap = tipCap
	auth.GasFeeCap = gasFeeCap
	auth.GasLimit = 300000         // Estimasi Gas Limit
	auth.Value = big.NewInt(90000) // value we pay
	auth.Nonce = big.NewInt(int64(nonceAt))

	contract, err := lendingContract.NewLendingContract(cAddr, c)
	assert.Nil(t, err)

	transaction, err := contract.RepayLoan(auth)
	assert.Nil(t, err, "Error during loan repayment, possibly insufficient funds")
	t.Logf("transaction hash %s", transaction.Hash())
	/*
		receipt, err := waitForTransactionReceipt(c, transaction.Hash())
		assert.Nil(t, err, "Failed to get transaction receipt")
		assert.Equal(t, receipt.Status, uint64(1), "Transaction failed on blockchain")
		t.Logf("Transaction confirmed in block %d", receipt.BlockNumber.Uint64())*/

}

/*
	func waitForTransactionReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
		ctx := context.Background()
		for {
			receipt, err := client.TransactionReceipt(ctx, txHash)
			if err == nil {
				return receipt, nil
			}
			if err.Error() == "not found" {
				time.Sleep(1 * time.Second) // Tunggu 1 detik sebelum mencoba lagi
				continue
			}
			return nil, err
		}
	}
*/
func ifErrLog(err error, msg string) {
	if err != nil {
		log.Fatalf("Error in %s: %v", msg, err)
	}
}
