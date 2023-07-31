# Proof of Marriage

This project is the result of the ETHGlobal Paris 2023 Hackathon.

> Find the project page in the [ETHGlobal showcase (here)](https://ethglobal.com/showcase/proof-of-marriage-4hxah).


## Project Description

The main goal of this project, is to focus on relationship status on the Internet, and use the Ethereum ecosystem to realize a Wedding smart-contract deployer, and a Dapp you can use during a wedding ceremony.

Why should I be married on the blockchain?

Some people want to be married in front of God, in front of the law, or just with their friends and family. Using the blockchain, with it's transparency and it's immutability is definitely a great way to be married. That's what we think!

## How it's Made

This project is made with 2 parts:

1 - A Go CLI tool, you can use to customize your custom wedding smart contract. - It uses Geth an Abigen (a Go-Ethereum tool) to automatically realize Go bindings with the Solidity smart contract. 
[CLI README](./cli/README.md)

2 - A React Dapp, using ethers to interact with our contract. We are using the Ethereum Attestation Service EAS (EAS) which is the perfect framework for our use case. EAS is a public good for creating, verifying, and revoking on/off-chain attestations, so we choose it to create Wedding attestation.
[DAPP README](./frontend/README.md)


Our DApp is offering a live chat, to allow you to interact with your guest during a wedding. We didn't have sufficient time during the ETHGlobal's hackathon to use either Push protocol or XMTP, so it's just a common chat in pure Solidity (gas are required to send messages).