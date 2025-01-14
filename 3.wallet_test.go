package go_etherium_learn

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"testing"
)

func TestGenerateETHWallet(t *testing.T) {

	//GENERATE PRIVATE KEY
	privateKey, _ := crypto.GenerateKey()
	fromECDSA := crypto.FromECDSA(privateKey)
	fmt.Println("PRIVATE KEY : ", hexutil.Encode(fromECDSA))

	//GENERATE PUBLIC KEY
	publicKey := privateKey.PublicKey
	fromECDSAPub := crypto.FromECDSAPub(&publicKey)
	fmt.Println("PUBLIC KEY : ", hexutil.Encode(fromECDSAPub))

	//ADDRESS (represent simply public key)
	hex := crypto.PubkeyToAddress(publicKey).Hex()
	fmt.Println("ADDRESS : ", hex)

}

// Simply encrypt privateKey use pass with symmetric cryptography algorithm, for access just input pass code, encrypt
// wallet called keyStore file
// For example:
// If your Trezor (hardware wallet) gets stolen, the thief cannot access your BTC without the passcode.
// This is because hardware wallets like Trezor require both physical access and the correct passcode to decrypt and use the private key.
func TestKeyStore(t *testing.T) {
	//ENCRYPTION:
	//	1. Generate Private Key -> 2. Input Password -> 3. Derive Key (KDF) ->
	//		4. Encrypt Private Key (AES) -> 5. Generate MAC ->
	//		6. Store Encrypted Wallet Data (Ciphertext, MAC, Salt, etc.)
	//
	//DECRYPTION:
	//	1. Input Password -> 2. Derive Key (KDF) -> 3. Validate MAC ->
	//		4. Decrypt Private Key (AES) -> 5. Use Private Key

	password := "jamal"
	/*keyStore := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	account, _ := keyStore.NewAccount(password)
	fmt.Println(account.Address)*/

	file, _ := os.ReadFile("./wallet/UTC--2025-01-14T11-02-41.636529300Z--0d518a4c445bbfad90c8382a051a91087d930253")
	decryptKey, _ := keystore.DecryptKey(file, password)
	pData := crypto.FromECDSA(decryptKey.PrivateKey)
	fmt.Println("PRIVATE KEY : ", hexutil.Encode(pData))

	//get PUBLIC KEY
	publicKey := crypto.FromECDSAPub(&decryptKey.PrivateKey.PublicKey)
	fmt.Println("PUBLIC KEY : ", hexutil.Encode(publicKey))

	//get Address (represent simply public key)
	address := crypto.PubkeyToAddress(decryptKey.PrivateKey.PublicKey).Hex()
	fmt.Println("ADDRESS : ", address)

}
