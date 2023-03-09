package main

import (
	// Standard
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	// Third Party
	"github.com/alexflint/go-arg"

	// Internal libs

	"github.com/Telept-xyz/vanity-crypto-address-factory/arweave"
	"github.com/Telept-xyz/vanity-crypto-address-factory/common"
	"github.com/Telept-xyz/vanity-crypto-address-factory/ethereum"
)

// GoReleaser
var (
	version = "dev"
	commit  = "untagged"
)

// go-arg
type args struct {
	Crypto  string `arg:"-c,--crypto" default:"" help:"choose from ethereum, arweave"`
	Workers int    `arg:"-w,--workers" default:"1" help:"Number of workers to spawn"`
	Count   int    `arg:"-n,--number" default:"1" help:"Number of wallets to generate"`
	Output  string `arg:"-o,--output" default:"./keyfiles" help:"Output directory"`
	Pattern string `arg:"positional,required" help:"Regex pattern to match the wallet address"`
}

// set go-arg version and commit from GoReleaser
func (args) Version() string {
	return fmt.Sprintf("wave %v (%v)", version, commit[:8])
}

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

func writeToFile(outDir string, crypto string, wallet common.Wallet) {

	// get keyfile as json byte slice
	keyfile, err := json.Marshal(wallet)
	errcheck(err)

	keyfilePath := filepath.Join(outDir, crypto+"-keyfile-"+wallet.Address+".json")

	// Check if output directory exists
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		// If not, create it
		derr := os.Mkdir(outDir, 0777)
		errcheck(derr)
	}
	// Write keyfile to file
	ioutil.WriteFile(keyfilePath, keyfile, 0666)
	fmt.Println("[EMIT] keyfile:", keyfilePath)
}

func worker(workerId int, pattern string, crypto common.Crypto, walletChan chan<- common.Wallet) {
	for {
		// Generate wallet
		wallet := crypto.GenerateWallet()
		walletAddress := wallet.Address

		// check if wallet address matches the provided pattern
		match, err := regexp.MatchString(pattern, walletAddress)
		errcheck(err)

		fmt.Printf("[WORKER%v] address: %v | match: %v]\n", workerId, wallet.Address, match)

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
func main() {
	// parse commandline arguments
	var args args
	arg.MustParse(&args)

	numWorkers := args.Workers                // Number of workers to spawn
	numWallets := args.Count                  // Number of wallets to generate
	vanityPattern := args.Pattern             // Regex pattern to match the wallet address
	outDir := filepath.Join(args.Output)      // Output directory
	walletChan := make(chan common.Wallet, 1) // Channel to get wallets from workers

	fmt.Println("Pattern:", "/"+vanityPattern+"/")
	fmt.Println("Outputs:", outDir)
	fmt.Println("Workers:", numWorkers)
	fmt.Println("Wallets:", numWallets)

	crypto := NewCrypto(args.Crypto)

	for n := 1; n <= numWallets; n++ {
		// spawn workers
		for w := 1; w <= numWorkers; w++ {
			go worker(w, vanityPattern, crypto, walletChan)
		}

		// get wallet from worker
		k := <-walletChan

		fmt.Println("[MATCH] address:", k.Address)
		writeToFile(outDir, args.Crypto, k)

	}
}
