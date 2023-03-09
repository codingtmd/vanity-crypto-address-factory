package common

// Interface for the factory
type Crypto interface {
	GenerateWallet() Wallet
}

type Wallet struct {
	Address    string
	PrivateKey string
	PublicKey  string
}
