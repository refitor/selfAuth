package contracts

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/refitor/web3/contracts/rscore"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func LoadContract(client *ethclient.Client, contractAddr string, fGetParam func(string) string) (*rscore.SelfAuthSession, error) {
	// Init new authenticated session
	session, err := doNewSession(client, fGetParam)
	if err != nil {
		return nil, err
	}
	err = doLoadContract(session, client, contractAddr)
	return session, err
}

// NewSession returns a quiz.SelfAuthSession struct that
// contains an authentication key to sign transactions with.
func doNewSession(client *ethclient.Client, fGetParam func(string) string) (*rscore.SelfAuthSession, error) {
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
	return &rscore.SelfAuthSession{
		TransactOpts: *auth,
		CallOpts: bind.CallOpts{
			From:    auth.From,
			Context: context.Background(),
		},
	}, nil
}

//// Contract initialization functions

// LoadContract loads a contract if one exists
func doLoadContract(session *rscore.SelfAuthSession, client *ethclient.Client, contractAddr string) error {
	if contractAddr == "" {
		return errors.New(("could not find a contract address to load"))
	}

	addr := common.HexToAddress(contractAddr)
	instance, err := rscore.NewSelfAuth(addr, client)
	if err != nil {
		return fmt.Errorf("could not load contract: %v\n", err)
	}
	session.Contract = instance
	return nil
}
