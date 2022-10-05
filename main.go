package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func queryBlockNumber(url string, nameofchain string) {
	client, err := ethclient.Dial(url) // get the latest block number
	if err != nil {
		log.Fatal(err)
	}
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var blockNumber_hex = strconv.FormatInt(int64(blockNumber), 16) //convert blocknumber to hex
	fmt.Println("##########################", nameofchain, "#####################################################################")
	fmt.Println("BLOCK NUMBER:", blockNumber)
	fmt.Println("BLOCK NUMBER IN HEX:", blockNumber_hex)
	fmt.Println("####################################################################################################################")
}

func Celo_APi() {
	const url = "https://explorer.celo.org/api?module=block&action=eth_block_number"
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
func main() {
	const url = "https://rpc.ankr.com/celo" // url string
	const nameofchain = "CELO"              // name of chain
	queryBlockNumber(url, nameofchain)
	Celo_APi()

}
