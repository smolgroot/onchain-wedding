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

