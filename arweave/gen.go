package arweave

import (
	// Standard
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"math/big"

	// Third Party
	"github.com/Telept-xyz/vanity-crypto-address-factory/common"
)

type Arweave struct{}

func (a Arweave) GenerateWallet() common.Wallet {
	// generate an RSA key
	reader := rand.Reader
	rsaKey, err := rsa.GenerateKey(reader, 4096)
	errcheck(err)

	// Generate wallet address
	h := sha256.New()
	h.Write(rsaKey.N.Bytes())                                   // Take the "n", in bytes and hash it using SHA256
	address := base64.RawURLEncoding.EncodeToString(h.Sum(nil)) // Then base64url encode it to get the wallet address

	return common.Wallet{Address: address, PrivateKey: base64.RawURLEncoding.EncodeToString(rsaKey.D.Bytes()), PublicKey: base64.RawURLEncoding.EncodeToString(big.NewInt(int64(rsaKey.E)).Bytes())}
}

func errcheck(e error) {
	if e != nil {
		panic(e)
	}
}
