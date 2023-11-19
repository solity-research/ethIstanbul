package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

type NetworkListener struct {
	url             string
	contractAddress string
}

type CreateEvent struct {
	Data        string
	NewContract common.Address
}
type MsgEvent struct {
	Message string
}

func (c *NetworkListener) listenNetwork() {
	client, err := ethclient.Dial(c.url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to wss")

	contractAddress := common.HexToAddress(c.contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	log.Println("filter is ready")

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening")
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Println("log hash", vLog.Topics[0])
			if vLog.Topics[0] == common.HexToHash("0x42da151aa690e925a014ed75bc20606fec61ebde4364808d7d5af5cb92efa813") {
				//Msg event triggered
				fmt.Println("Topic", vLog.Topics)
				fmt.Println("Data", vLog.Data)
				fmt.Println("Adresss", vLog.Address)
				abiFile, err := ioutil.ReadFile("contract.abi")
				contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
				if err != nil {
					// Handle error
					fmt.Println(err)
				}
				event := MsgEvent{}
				err = contractAbi.UnpackIntoInterface(&event, "PromptSent", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// 0=>message
				fmt.Println("Event", event)
			} else if vLog.Topics[0] == common.HexToHash("0x0d1802b86a0633c4679107d4313baeacab2fb7391348ca6485df3f40844b0b07") {
				//Create event triggered
				fmt.Println("Topic", vLog.Topics)
				fmt.Println("Data", vLog.Data)
				fmt.Println("Adresss", vLog.Address)
				abiFile, err := ioutil.ReadFile("contract.abi")
				contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
				if err != nil {
					// Handle error
					fmt.Println(err)
				}
				event := CreateEvent{}
				err = contractAbi.UnpackIntoInterface(&event, "ContractCreated", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// 0=>Data 1=>NewContract
				fmt.Println("Event", event)
			}

		}
	}
}

func (c *NetworkListener) sendCartesiCreate(data CreateEvent) {
	client, err := ethclient.Dial(c.url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to localhost")
	contractAddress := common.HexToAddress(c.contractAddress)

}

func (c *NetworkListener) sendCartesiMessage(msg MsgEvent) {
	client, err := ethclient.Dial(c.url)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to localhost")
	contractAddress := common.HexToAddress(c.contractAddress)
}

func CreateFunctionRequirementsForLendingPool(clientUrl string, lendingPoolAddress string, privateKey string) (error, *Bridge, common.Address, *bind.TransactOpts) {
	err, client, _publicAddress, res := CreateFunctionRequirementsForControllers(
		clientUrl,
		"LendingPool.abi",
		lendingPoolAddress,
		privateKey)

	address := common.HexToAddress(lendingPoolAddress)
	contractInstance, err := NewBridge(address, client)
	return err, contractInstance, _publicAddress, res
}

func CreateFunctionRequirementsForControllers(clientUrl string, walletAbiName string, oracleAddress string, privateKey string) (error, *ethclient.Client, common.Address, *bind.TransactOpts) {
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		// Handle error
	}

	address := common.HexToAddress(oracleAddress)
	abiFile, err := ioutil.ReadFile(walletAbiName)
	_, err = abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	fmt.Println(address)
	//fmt.Println(client)
	//fmt.Println(contractAbi)

	if err != nil {
		// Handle error
	}

	_privateKey, _, _publicAddress, _ := GenerateKeypairFromPrivateKeyHex(privateKey)
	res, _ := BuildTransactionOptions(client, _publicAddress, _privateKey, 300000)
	return err, client, _publicAddress, res
}

func BuildTransactionOptions(client *ethclient.Client, fromAddress common.Address, prvKey *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {

	// Retrieve the chainID
	chainID, cIDErr := client.ChainID(context.Background())

	if cIDErr != nil {
		return nil, cIDErr
	}

	// Retrieve suggested gas price
	gasPrice, gErr := client.SuggestGasPrice(context.Background())

	if gErr != nil {
		return nil, gErr
	}

	// Create empty options object
	txOptions, txOptErr := bind.NewKeyedTransactorWithChainID(prvKey, chainID)

	if txOptErr != nil {
		return nil, txOptErr
	}

	txOptions.Value = big.NewInt(0) // in wei
	txOptions.GasLimit = gasLimit   // in units
	txOptions.GasPrice = gasPrice

	return txOptions, nil
}

func GenerateKeypairFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address, error) {

	log.Println("Generating the keypair...")

	// If hex string has "0x" at the start discard it
	if privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Convert hex key to a private key object
	privateKey, privateKeyErr := crypto.HexToECDSA(privateKeyHex)

	if privateKeyErr != nil {
		return nil, nil, common.Address{}, privateKeyErr
	}

	// Generate the public key using the private key
	publicKey := privateKey.Public()

	// Cast crypto.Publickey object to the ecdsa.PublicKey object
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}

	// Convert publickey to a address
	pubkeyAsAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, pubkeyAsAddress, nil
}
func main() {
	var networkListener = NetworkListener{
		url:             "wss://omniscient-newest-log.scroll-testnet.quiknode.pro/d523803d9f279c6ae232b4a48953cbb8477a760b/",
		contractAddress: "0x75A14E9109eDBa761CA8f5F3A5ea662fd28E3546",
	}
	for {
		networkListener.listenNetwork()
	}
}
