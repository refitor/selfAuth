package web3

import (
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var contractABI abi.ABI
var getParamFunc func(string) string

type SelfAuthEvent struct {
	AuthAddr [20]byte
	Params   []byte
}

var v_contract_abi_ISelfAuth = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"authAddr","type":"address"},{"indexed":false,"internalType":"bytes","name":"params","type":"bytes"}],"name":"authRequest","type":"event"},{"inputs":[{"internalType":"address","name":"authAddr","type":"address"},{"internalType":"bytes","name":"params","type":"bytes"}],"name":"authResponse","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

func Init(gatewayWss, gatewayHttps string, fGetParam func(string) string) (*ethclient.Client, error) {
	contractAbi, err := abi.JSON(strings.NewReader(v_contract_abi_ISelfAuth))
	if err != nil {
		log.Fatal(err)
	}
	contractABI = contractAbi
	getParamFunc = fGetParam
	return ethclient.Dial(gatewayWss)
}

func UnInit(client *ethclient.Client) {
	if client != nil {
		client.Close()
	}
}

func WatchTransaction(client *ethclient.Client, watchAddress, to string) (err error) {
	err = doWatchContract(client, watchAddress, func(l types.Log) error {
		return nil
	}, true)
	return
}

func WatchContract(client *ethclient.Client, watchAddress string, handleFunc func(types.Log) error) {
	go doWatchContract(client, watchAddress, handleFunc, false)
}

func doWatchContract(client *ethclient.Client, watchAddress string, handleFunc func(types.Log) error, bOnce bool) error {
	contractAddress := common.HexToAddress(watchAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:

			// fmt.Printf("%+v\n", vLog) // pointer to event log

			if handleFunc != nil {
				if bOnce {
					return handleFunc(vLog)
				} else {
					go handleFunc(vLog)
				}
			}
		}
	}
}

// stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func StringToKeccak256(s string) [32]byte {
	var output [32]byte
	copy(output[:], crypto.Keccak256([]byte(s))[:])
	return output
}

func GetEvent(l types.Log) *SelfAuthEvent {
	event := &SelfAuthEvent{}
	err := contractABI.UnpackIntoInterface(event, "authRequest", l.Data)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return event
}

func GetEventAddress(e *SelfAuthEvent) string {
	return common.BytesToAddress(e.AuthAddr[:]).String()
}
