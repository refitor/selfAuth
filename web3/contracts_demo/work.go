package contracts_demo

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var v_contract_address string

func ContractAddress() string {
	return v_contract_address
}

func Init(ctx context.Context, client *ethclient.Client, fGetParam func(string) string, contractDatas ...interface{}) (*SADemoSession, error) {
	// defer client.Close()
	v_contract_address = fGetParam("CONTRACTADDR")

	// Init new authenticated session
	session, err := doNewSession(ctx, client, fGetParam)
	if err != nil {
		return nil, err
	}

	// Load or Deploy contract, and update session with contract instance
	if v_contract_address == "" && fGetParam("operate") == "CONTRACTNEW" {
		err = doNewContract(session, client, contractDatas...)
	}

	// If we have an existing contract, load it; if we've deployed a new contract, attempt to load it.
	if v_contract_address != "" {
		err = doLoadContract(session, client)
	}
	return session, err
}

// NewSession returns a quiz.SADemoSession struct that
// contains an authentication key to sign transactions with.
func doNewSession(ctx context.Context, client *ethclient.Client, fGetParam func(string) string) (*SADemoSession, error) {
	// Create new transactor
	keystore, err := os.Open(fGetParam("KEYSTORE"))
	if err != nil {
		return nil, err
	}
	defer keystore.Close()

	auth, err := bind.NewTransactor(keystore, fGetParam("KEYSTOREPASS"))
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	// bind.NewTransactor() returns a bind.TransactOpts{} struct with the following field values:
	// From: auth.From,
	// Signer: auth.Signer,
	// Nonce: nil // Setting to nil uses nonce of pending state
	// Value: big.NewInt(0), // 0 because we're not transferring Eth
	// GasPrice: nil // Setting to nil automatically suggests a gas price
	// GasLimit: 0 // Setting to 0 automatically estimates gas limit

	// Return session without contract instance
	return &SADemoSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: ctx,
		},
	}, nil
}

//// Contract initialization functions

// NewContract deploys a contract if no existing contract exists
func doNewContract(session *SADemoSession, client *ethclient.Client, contractDatas ...interface{}) error {
	// Test our inputs
	if v_contract_address != "" {
		return nil
	}

	// Hash answer before sending it over Ethereum network.
	contractAddress, tx, instance, err := DeploySADemo(&session.TransactOpts, client)
	if err != nil {
		return fmt.Errorf("could not deploy contract: %v\n", err)
	}
	fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

	session.Contract = instance
	v_contract_address = contractAddress.Hex()
	return nil
}

// LoadContract loads a contract if one exists
func doLoadContract(session *SADemoSession, client *ethclient.Client) error {
	if v_contract_address == "" {
		return errors.New(("could not find a contract address to load"))
	}

	addr := common.HexToAddress(v_contract_address)
	instance, err := NewSADemo(addr, client)
	if err != nil {
		return fmt.Errorf("could not load contract: %v\n", err)
	}
	session.Contract = instance
	return nil
}

// stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func stringToKeccak256(s string) [32]byte {
	var output [32]byte
	copy(output[:], crypto.Keccak256([]byte(s))[:])
	return output
}
