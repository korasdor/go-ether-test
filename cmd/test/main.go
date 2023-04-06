package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// infuraUrl  = "https://mainnet.infura.io/v3/17ec44fd8ad4457cafb926b75e360484"
	alchemyUrl = "https://eth-mainnet.g.alchemy.com/v2/aeyWf4ZgLuvyQKE7ac1rewRSi4CZVPt4"
)

func main() {
	//create client
	client, err := ethclient.DialContext(context.Background(), alchemyUrl)
	if err != nil {
		log.Fatalf("Error to create a ether client %v", err)
	}
	defer client.Close()

	//get last block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block %v", err)
	}
	fmt.Println("The block number:", block.Number())

	//get balance
	addrHex := "0xE94f1fa4F27D9d288FFeA234bB62E1fBC086CA0c"
	addr := common.HexToAddress(addrHex)

	balance, err := client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		log.Fatalf("Error to get a balance %v", err)
	}
	fmt.Println("The balance:", balance)

	//convert balance to ehter (1 ehter = 10^18 wei)
	fBalane := &big.Float{}
	fBalane.SetString(balance.String())

	balanceEther := &big.Float{}
	balanceEther.Quo(fBalane, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceEther)
}
