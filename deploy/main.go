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
	auth.GasLimit = 5000000    // Gas limit for executing the contract or transaction (varies by case)
	auth.GasFeeCap = gasFeeCap // Maximum gas fee we're willing to pay
	auth.GasTipCap = tipCap    // Priority fee (tip) for faster transaction processing

	balance, err := c.BalanceAt(ctx, crypto.PubkeyToAddress(pubKey), nil)
	ifErrLog(err, "fetching balance")
	fmt.Printf("Balance: %s wei\n", balance.String())

	// deploy contract
	interestRate := big.NewInt(5000)
	loanAmountLimit := big.NewInt(10000000000)
	contract, transaction, _, err := lendingContract.DeployLendingContract(auth, c, interestRate, loanAmountLimit)
	ifErrLog(err, "deploying smart contract")

	// Log alamat kontrak dan transaksi deploy
	fmt.Printf("Contract deployed at: %s\n", contract.Hex())
	fmt.Printf("Transaction hash: %s\n", transaction.Hash().Hex())
	fmt.Printf("Gas Limit: %d, Gas Fee Cap: %s, Gas Tip Cap: %s\n", auth.GasLimit, auth.GasFeeCap.String(), auth.GasTipCap.String())

	// Verifikasi apakah kode kontrak ada di blockchain
	fmt.Println("Verifying contract code on blockchain...")
	code, err := c.CodeAt(ctx, contract, nil)
	ifErrLog(err, "fetching contract code")
	if len(code) == 0 {
		log.Fatalf("No contract code found at address: %s\n", contract.Hex())
	}
	fmt.Println("Contract code successfully found on blockchain!")

	fmt.Printf("%s SUMMARY %s\n", strings.Repeat("-", 35), strings.Repeat("-", 35))
	fmt.Printf("CONTRACT        <- %s\n", contract)
	fmt.Printf("Tx              <- %v\n", transaction)
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

/*
cek contract int -> https://sepolia.etherscan.io/address/0xB2528bD8E59988E70a849e9eedd837C3c598907d
----------------------------------- SUMMARY -----------------------------------
CONTRACT        <- 0xB2528bD8E59988E70a849e9eedd837C3c598907d
Tx              <- &{0xc00019e980 {13968622154990673676 2513358501 0x7ff6b5813d60} {[] {} <nil>} {{} {} 0} {[] {} <nil>}}
PRIVATE KEY OWN <- d95268fd0b131fc1b8c34e5c402236cd4319a739098e38bd156273ce134cfcb0
PUBLIC KEY OWN  <- 04165d4764cc62cf72bb58f547a52742de7998b612823028cab87722fd480c1d04e35e94a8953d01d9d45a516a31820854e39cea279295448b7413037309086326
ADDR OWN        <- 0x0d518a4C445bBfaD90C8382a051A91087d930253
----------------------------------------------------------------------
*/
