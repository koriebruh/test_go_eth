package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	lendingContract "koriebruh/learn/gen"
	"log"
	"math/big"
	"os"
	"strings"
)

func main() {
	var ctx = context.Background()
	var infraURL = "https://sepolia.infura.io/v3/7108f6b019944d2082df7b667e6b1f4a"
	//var ganacheURL = "http://127.0.0.1:8545"
	psw := "jamal"

	file, err := os.ReadFile("wallet/UTC--2025-01-14T11-02-41.636529300Z--0d518a4c445bbfad90c8382a051a91087d930253")
	ifErrLog(err, "reading wallet file")

	// GET OWNER DATA FOR OWN SMART CONTRACT
	decrypt, err := keystore.DecryptKey(file, psw)
	ifErrLog(err, "decrypting wallet file")
	pvKey := decrypt.PrivateKey
	pvKeyBytes := crypto.FromECDSA(pvKey)
	pubKey := decrypt.PrivateKey.PublicKey
	pubKeyBytes := crypto.FromECDSAPub(&pubKey)
	addr := crypto.PubkeyToAddress(pubKey).Hex()
	fmt.Printf("ADDR OWN        <- %s\n", addr)

	// CONNECT CLIENT
	c, err := ethclient.DialContext(ctx, infraURL)
	defer c.Close()
	ifErrLog(err, "connecting to Ethereum client")

	chainID, err := c.ChainID(ctx)
	ifErrLog(err, "fetching chain ID")
	nonceAt, err := c.PendingNonceAt(ctx, crypto.PubkeyToAddress(pubKey))
	ifErrLog(err, "fetching pending nonce")
	//tipCap, err := c.SuggestGasTipCap(ctx)
	//ifErrLog(err, "fetching suggested gas tip cap")

	baseFee := big.NewInt(20000000)                // Base fee, misalnya 20 Gwei
	tipCap := big.NewInt(1000000000)               // Tip fee, 1 Gwei
	gasFeeCap := new(big.Int).Add(baseFee, tipCap) // Total max fee

	auth, err := bind.NewKeyedTransactorWithChainID(pvKey, chainID)
	ifErrLog(err, "creating transaction signer")
	auth.Nonce = big.NewInt(int64(nonceAt))
	auth.GasLimit = 1000000    // Gas limit for executing the contract or transaction (varies by case)
	auth.GasFeeCap = gasFeeCap // Maximum gas fee we're willing to pay
	auth.GasTipCap = tipCap    // Priority fee (tip) for faster transaction processing

	balance, err := c.BalanceAt(ctx, crypto.PubkeyToAddress(pubKey), nil)
	ifErrLog(err, "fetching balance")
	fmt.Printf("Balance: %s wei\n", balance.String())

	interestRate := big.NewInt(5000)
	loanAmountLimit := big.NewInt(10000000000)
	contract, transaction, _, err := lendingContract.DeployLendingContract(auth, c, interestRate, loanAmountLimit)
	ifErrLog(err, "deploying smart contract")

	fmt.Printf("%s SUMMARY %s\n", strings.Repeat("-", 35), strings.Repeat("-", 35))
	fmt.Printf("CONTRACT        <- %x\n", contract)
	fmt.Printf("Tx              <- %x\n", transaction)
	fmt.Printf("PRIVATE KEY OWN <- %s\n", hex.EncodeToString(pvKeyBytes))
	fmt.Printf("PUBLIC KEY OWN  <- %s\n", hex.EncodeToString(pubKeyBytes))
	fmt.Printf("ADDR OWN        <- %s\n", addr)
	fmt.Printf("%s", strings.Repeat("-", 70))
}

func ifErrLog(err error, context string) {
	if err != nil {
		log.Fatalf("Error in %s: %v", context, err)
	}
}
