/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Run this command to configure your wedding smart contract with a few interactive steps.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n🎂✨💍 WELCOME IN THE CONFIGURATION MENU OF YOUR WEDDING SMART CONTRACT 💒🕺💖\n")

		// os.Setenv("WEDDING_SMART_CONTRACT_CONFIGURED", "true")
		checkedState := os.Getenv("WEDDING_SMART_CONTRACT_CONFIGURED")
		if checkedState != "true" {
			configure()
		} else {
			fmt.Println("Smart contract seems to be already configured")
			fmt.Println("You can now deploy it on the blockchain. Please check the `config.json` for the current ChainID and RPC address.")
			fmt.Println("\nUsage:")
			fmt.Println("\tsmartWedding deploy")
		}

	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

type WeddingsDetails struct {
	Addr1  string
	Addr2  string
	Name1  string
	Name2  string
	Descr1 string
	Descr2 string
	Ipfs1  string
	Ipfs2  string
}

func configure() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("[Fiance 1] What's the first partner firstname?\t")
	name1, _ := reader.ReadString('\n')
	name1 = strings.TrimSuffix(name1, "\n")

	fmt.Print("[Fiance 1] What's the first partner Ethereum address?\t")
	addr1, _ := reader.ReadString('\n')
	addr1 = strings.TrimSuffix(addr1, "\n")

	fmt.Print("[Fiance 1] Describe the first partner in a short sentence.\t")
	descr1, _ := reader.ReadString('\n')
	descr1 = strings.TrimSuffix(descr1, "\n")

	fmt.Print("[Fiance 1] Give a photo's IPFS cid.\t")
	ipfs1, _ := reader.ReadString('\n')
	ipfs1 = strings.TrimSuffix(ipfs1, "\n")

	fmt.Print("[Fiance 2] What's the second partner firstname?\t")
	name2, _ := reader.ReadString('\n')
	name2 = strings.TrimSuffix(name2, "\n")

	fmt.Print("[Fiance 2] What's the second partner Ethereum address?\t")
	addr2, _ := reader.ReadString('\n')
	addr2 = strings.TrimSuffix(addr2, "\n")

	fmt.Print("[Fiance 2] Describe the second partner in a short sentence.\t")
	descr2, _ := reader.ReadString('\n')
	descr2 = strings.TrimSuffix(descr2, "\n")

	fmt.Print("[Fiance 2] Give a photo's IPFS cid.\t")
	ipfs2, _ := reader.ReadString('\n')
	ipfs2 = strings.TrimSuffix(ipfs2, "\n")
	fmt.Printf(ipfs2)

	data := WeddingsDetails{
		Addr1:  addr1,
		Addr2:  addr2,
		Name1:  name1,
		Name2:  name2,
		Descr1: descr1,
		Descr2: descr2,
		Ipfs1:  ipfs1,
		Ipfs2:  ipfs2,
	}
	tmpl := template.New("WeddingTmpl")

	// tmpl, _ = tmpl.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")
	tmpl, err := template.ParseFiles("Wedding.tmpl")
	if err != nil {
		_ = fmt.Errorf("cannot parse smart contract template")
		panic(err)
	}

	// // standard output to print merged data
	// err1 := tmpl.Execute(os.Stdout, data)
	// if err1 != nil {
	// 	panic(err1)
	// }

	// // if there is no error,
	// // prints the output
	// if err == nil {
	// 	fmt.Println(err)
	// }

	// Write to file
	f, err := os.Create("./contract/Wedding.sol")
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}
	err2 := tmpl.Execute(f, data)
	if err2 != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()

	// Congrats! Here we go!
	fmt.Println("\nCongrats!")
	fmt.Println("Next step: run the following command now")
	fmt.Println("\n\t$ ./build.sh")
	fmt.Println("\t$ smartWedding deploy")

}
