package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery" //scraping data from html
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
	const url2 = "https://explorer.celo.org/api?module=token&action=bridgedTokenList"
	const url3 = "https://explorer.celo.org/blocks"

	response, err := http.Get(url)
	response2, err := http.Get(url2)
	response3, err := http.Get(url3)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()
	defer response2.Body.Close()
	defer response3.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	content2, _ := ioutil.ReadAll(response2.Body)
	//content3, _ := ioutil.ReadAll(response3.Body)
	fmt.Println(string(content))
	fmt.Println(string(content2))
	//fmt.Println(string(content3))
	doc, err := goquery.NewDocumentFromReader(response3.Body)
	if err != nil {
		log.Fatal(err)
	}
	title_size := doc.Find("div.card").Size() //find the size of the div
	title_size2 := doc.Find("div.tile").Size()
	fmt.Println("SIZE OF THE DIV:", title_size)
	fmt.Println("SIZE OF THE DIV:", title_size2)
	title, error := doc.Find("div.card").Html()
	title2, error := doc.Find("div.tile tile-type-block fade-up").Html()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("HTML:", title)
	fmt.Println("HTML:", title2)
}
func main() {
	const url = "https://rpc.ankr.com/celo" // url string
	const nameofchain = "CELO"              // name of chain
	queryBlockNumber(url, nameofchain)
	Celo_APi()
	//geth --datadir DATADIR

}
