package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/Telept-xyz/vanity-crypto-address-factory/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Structs for different shapes
type Ethereum struct{}

func (e Ethereum) GenerateWallet() common.Wallet {
	// Generate a private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	// Convert private key to hex string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	//fmt.Println("Private key:", privateKeyHex)

	// Get the Ethereum address from the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("failed to cast public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println("Address:", address)

	// Convert the public key to a string
	publicKeyString := hex.EncodeToString(crypto.FromECDSAPub(publicKeyECDSA))

	return common.Wallet{Address: address, PrivateKey: privateKeyHex, PublicKey: publicKeyString}
}
