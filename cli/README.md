# Requirements

1. You need to install Go Ethereum Tools (both `geth` and `abigen` are required).
Please visit https://geth.ethereum.org/downloads to choose the release the suits your environment.

2. You also need to get the `solc` Solidity compiler. You can install it running
    ```bash
    npm install -g solc@0.8.16
    ```

# Get Ready

Compile de Solidity Smart Contract using `solcjs`

```bash
solcjs contract/SmartWedding.sol --abi --bin -o contract  
```

Auto generate Go bindings to interact with the smart contract within our CLI

```bash
abigen --abi contract/SmartWedding_sol_SmartWedding.abi --bin contract/SmartWedding_sol_SmartWedding.bin --pkg utils --type SmartWedding --out utils/bindings.go
```


# Compile the CLI

```bash
go build .
./smartwedding --help
```

