package contract

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	ChainAPIURL = "https://goerli.infura.io/v3/47aef1ccbb464b75af6d2b427cb7b8b4" //https://mainnet.infura.io/v3/47aef1ccbb464b75af6d2b427cb7b8b4
)

func EthTxCli(stampAddr string, pri *ecdsa.PrivateKey) (*StampCaller, *bind.TransactOpts, error) {

	cli, err := ethclient.Dial(ChainAPIURL)
	if err != nil {
		return nil, nil, err
	}

	ctr, err := NewStampCaller(common.HexToAddress(stampAddr), cli)
	if err != nil {
		return nil, nil, err
	}

	chainId, err := cli.ChainID(context.TODO())
	if err != nil {
		panic(err)
	}

	transOpt, err := bind.NewKeyedTransactorWithChainID(pri, chainId)
	if err != nil {
		return nil, nil, err
	}
	return ctr, transOpt, nil
}

func StampBalanceOfWallet(addr common.Address) (*big.Int, error) {
	cli, err := ethclient.Dial(ChainAPIURL)
	if err != nil {
		return nil, err
	}

	ctr, err := NewStampCaller(addr, cli)
	if err != nil {
		return nil, err
	}

	return ctr.BalanceOf(nil, addr)
}
