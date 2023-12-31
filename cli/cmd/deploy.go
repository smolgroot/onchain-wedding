/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"smartWedding/utils"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a new instance of the Wedding smart contract on chain. Returns contract address.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		deploy()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deploy() {

	config := utils.LoadConfiguration("./config.json")

	// Connect to our custom RPC
	conn, err := ethclient.Dial(config.RPCAddress)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum RPC: %v", err)
	}

	// Unlock the wallet
	fmt.Print("Wallet Password: ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	pass := string(bytepw)

	// Init a transactor for our chain
	// Give the wallet password to sign the tx
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(config.KeyStore), pass, big.NewInt(config.ChainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set the gas price
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, tx, _, err := utils.DeployWedding(auth, conn)

	if err != nil {
		log.Fatalf("Failed to deploy new task contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// time.Sleep(10 * time.Second) // Allow it to be processed by the local node :P

}
