#!/bin/bash

# set -e

# clean old contract metadata
rm contract/*.abi
rm contract/*.bin

# Compile the contract
solcjs contract/Wedding.sol --abi --bin -o contract/

# Rename the files to keep it easy
mv contract/*.abi contract/wedding.abi
mv contract/*.bin contract/wedding.bin

abigen --abi contract/wedding.abi --bin contract/wedding.bin --pkg utils --type Wedding --out utils/bindings.go

# Build the GO CLI
go build .