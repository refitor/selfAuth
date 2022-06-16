package web3

import (
	"github.com/refitor/selfAuth/web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AuthResponse(client *ethclient.Client, contractAddr, authAddr string, params []byte) (string, error) {
	session, err := contracts.LoadContract(client, contractAddr, getParamFunc)
	if err != nil {
		return "", err
	}

	tx, err := session.AuthResponse(common.HexToAddress(authAddr), params)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), err
}
