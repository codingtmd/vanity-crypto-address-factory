package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

func search(prefix string, postfix string) {
	for {
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
		fmt.Println("Address:", address)

		// Check if the address starts with a certain substring
		if address[:len(prefix)] == prefix && address[len(address)-len(postfix):] == postfix {
			fmt.Printf("Find it!!!!!Address starts with %s\n", prefix)

			writeToFile(fmt.Sprintf("private key %s, address %s \n", privateKeyHex, address))
			break
		}
	}
}

func writeToFile(str string) {
	out, e := os.Create("vanity.txt")
	if e != nil {
		panic(e)
	}
	defer out.Close()
	out.WriteString(str)
}

func main() {
	search("0x666", "8888")
}
