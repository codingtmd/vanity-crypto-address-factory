# VAS: vanity-eth-address

## Why this repo
I saw quite a lot beautiful address in the web3, which let me thought, it might be a common needs for everyone. So I decide to write a helper for anyone to find the crypto address meaningful.

What's a vanity address? A vanity address is an address which part of it is chosen by yourself, making it look less random.

Examples: 0xc0ffee254729296a45a3885639AC7E10F8888888, or 0x999999cf1046e68e36E1aA2E0E07105eDDD1f08E. the previous one is ending with 8888888, the last one ie starting with 999999. You can also define any pattern for yourself.


The mechanism behinds the code is really simple. it is similiar as collison attach. The code will keep generating address UNTIL there is one matching the regex pattern. So, if your patter is complicated, it might take couple of days for the code to find out your dream address. 

Hope you like this tool and it helps you.


## Security
As you can check the code by youself, everything is computed only in your localbox. Nothing ever leaves your machine, or even your command line tool. Everything vanishes when you close your runtime window.

## Compatibility
Any address generated with Vanity-ETH is compatible with the crypto type you choose, which means you can use it for an ICO, an airdrop, or just to withdraw your funds from an exchange.

Normally the wallet (no matter MyEtherWallet, MetaMask, Mist, Arweave or whatever) will provide a way as "import your wallet with private key". And you can start from there

## I'm a developer, what I can help
Ok, I think two things will be really helpful for this tool
1. add more supported crypto type. As you see, it only support ethereum and arweave. Welcome to submit PR for any new crypto. For any newcrypto adding, plz create a separated folder and code it following the pattern
2. add opencl GPU accelaration. This code is majorly on cpu now, whch means SLOW. And unfortunately, I'm not farmiliar with GPU coding. If any expert can help on, it will be awesome to accelerate everyone's treasure-huntings



## Build from source
```
go mod tidy
go build
go run main.go -c "arweave" "^6.*8$"
```

## Usage

```
vca -h
wave 0.2.0 (8e7c6876)
Usage: vca [--workers WORKERS] [--number NUMBER] [--output OUTPUT] PATTERN

Positional arguments:
  PATTERN                Regex pattern to match the wallet address

Options:
  --workers WORKERS, -w WORKERS
                         Number of workers to spawn [default: 4]
  --number NUMBER, -n NUMBER
                         Number of wallets to generate [default: 1]
  --output OUTPUT, -o OUTPUT
                         Output directory [default: ./keyfiles]
  --help, -h             display this help and exit
  --version              display version and exit
```

### Example

```
vca -c "ethereum" ".*8$"
Pattern: /.*8$/
Outputs: keyfiles
Workers: 1
Wallets: 1
[WORKER1] address: 0xA1aD3315c2fcbf1d5BCa590a4F28A30406E71760 | match: false]
[WORKER1] address: 0x9b5cDc04FE34C0BECd04f5C048a9FD8ac3aa8005 | match: false]
[WORKER1] address: 0x9B2d82f720E442946907938917FF5efa26295ff6 | match: false]
[WORKER1] address: 0xd549A4910756526CfF3bcb1B8EEcb91e1784152b | match: false]
[WORKER1] address: 0xeB4Ac11100Cf5c895144bA6E7c16F0DE6cb2C6af | match: false]
[WORKER1] address: 0xC94A7f5470E3bE4919F882cD3Ae0F7D76ce4A8E8 | match: true]
[MATCH] address: 0xC94A7f5470E3bE4919F882cD3Ae0F7D76ce4A8E8
[EMIT] keyfile: keyfiles\ethereum-keyfile-0xC94A7f5470E3bE4919F882cD3Ae0F7D76ce4A8E8.json
```

## Reference
Thanks https://github.com/maximousblk/wave.go for the arweave code


## Tips
If you like the tool, you can send your tips over to `0x90822EE56ffBC14F3216846D967459aF268ed80C` 💛💛💛. ERC20 token is accepted.