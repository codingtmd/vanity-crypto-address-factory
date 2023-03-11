package main

import (
	// Standard
	"fmt"
	"regexp"
	"syscall/js"

	// Third Party

	// Internal libs

	"github.com/Telept-xyz/vanity-crypto-address-factory/arweave"
	"github.com/Telept-xyz/vanity-crypto-address-factory/common"
	"github.com/Telept-xyz/vanity-crypto-address-factory/ethereum"
)

// Factory function to create crypto type
func NewCrypto(cryptoType string) common.Crypto {
	switch cryptoType {
	case "ethereum":
		return ethereum.Ethereum{}
	case "arweave":
		return arweave.Arweave{}
	default:
		return nil
	}
}

func toJs(key string, value interface{}) {
	alert := js.Global().Get("onVcafMsg")
	alert.Invoke(key, value)
}

func worker(workerId int, pattern string, crypto common.Crypto, walletChan chan<- common.Wallet) {
	tried := 0
	for {
		// Generate wallet
		wallet := crypto.GenerateWallet()
		walletAddress := wallet.Address

		// check if wallet address matches the provided pattern
		match, err := regexp.MatchString(pattern, walletAddress)
		errcheck(err)

		// fmt.Printf("[WORKER%v] address: %v | match: %v]\n", workerId, wallet.Address, match)

		tried++

		if tried%10 == 0 {
			toJs("worker", map[string]interface{}{
				"workerId": workerId,
				"tried":    tried,
				"address":  wallet.Address,
			})
		}

		// send wallet to main if matched
		if match {
			walletChan <- wallet
			break
		}
	}
}

func errcheck(e error) {
	if e != nil {
		panic(e)
	}
}

/*
some example

	// find an address start with 6 and end with 8, you can change to any word you want
	go run main.go -c "arweave" "^6.*8$"
*/
func generate(cryptoType string, vanityPattern string, numWorkers int, numWallets int) bool {
	walletChan := make(chan common.Wallet, 1) // Channel to get wallets from workers

	fmt.Println("Pattern:", "/"+vanityPattern+"/")
	fmt.Println("Workers:", numWorkers)
	fmt.Println("Wallets:", numWallets)

	toJs("start", map[string]interface{}{})

	found := false

	crypto := NewCrypto(cryptoType)

	for n := 1; n <= numWallets; n++ {
		// spawn workers
		for w := 1; w <= numWorkers; w++ {
			go worker(w, vanityPattern, crypto, walletChan)
		}

		// get wallet from worker
		k := <-walletChan

		fmt.Println("[MATCH] address:", k.Address)

		toJs("found", map[string]interface{}{
			"address":    k.Address,
			"privateKey": k.PrivateKey,
			"publicKey":  k.PublicKey,
		})

		found = true
	}

	toJs("finish", map[string]interface{}{})

	return found
}

func generateFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(generate(args[0].String(), args[1].String(), args[2].Int(), args[3].Int()))
}

func main() {
	done := make(chan int, 0)
	js.Global().Set("vcafGenerate", js.FuncOf(generateFunc))
	<-done
}
