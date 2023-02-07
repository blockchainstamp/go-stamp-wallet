package contract

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

const (
	ChainAPIURL = "https://goerli.infura.io/v3/47aef1ccbb464b75af6d2b427cb7b8b4" //https://mainnet.infura.io/v3/47aef1ccbb464b75af6d2b427cb7b8b4
)

var (
	EmptyRecord = StampRecord{
		Value: big.NewInt(0),
		Nonce: big.NewInt(0),
	}
	_scLog = logrus.WithFields(logrus.Fields{
		"mode": "smart contract",
	})
)

func EthTxCli(stampAddr string, pri *ecdsa.PrivateKey) (*Stamp, *bind.TransactOpts, error) {

	cli, err := ethclient.Dial(ChainAPIURL)
	if err != nil {
		_scLog.Warn(err)
		return nil, nil, err
	}

	ctr, err := NewStamp(common.HexToAddress(stampAddr), cli)
	if err != nil {
		_scLog.Warn(err)
		return nil, nil, err
	}

	chainId, err := cli.ChainID(context.TODO())
	if err != nil {
		_scLog.Warn(err)
		return nil, nil, err
	}

	transOpt, err := bind.NewKeyedTransactorWithChainID(pri, chainId)
	if err != nil {
		_scLog.Warn(err)
		return nil, nil, err
	}
	return ctr, transOpt, nil
}

func StampBalanceOfWallet(stampAddr, userAddr common.Address) (StampRecord, error) {

	cli, err := ethclient.Dial(ChainAPIURL)
	if err != nil {
		_scLog.Warn(err)
		return EmptyRecord, err
	}

	ctr, err := NewStampCaller(stampAddr, cli)
	if err != nil {
		_scLog.Warn(err)
		return EmptyRecord, err
	}

	return ctr.BalanceOf(nil, userAddr)
}

func StampConfigFromBlockChain(addr string) (*StampConf, error) {

	if !common.IsHexAddress(addr) {
		return nil, fmt.Errorf("[%s] is invalid blockchain address", addr)
	}

	cli, err := ethclient.Dial(ChainAPIURL)
	if err != nil {
		_scLog.Warn(err)
		return nil, err
	}

	ctr, err := NewStampCaller(common.HexToAddress(addr), cli)
	if err != nil {
		_scLog.Warn(err)
		return nil, err
	}
	cf, err := ctr.Conf(nil)
	if err != nil {
		_scLog.Warn(err)
		return nil, err
	}
	return &StampConf{
		Addr:         addr,
		MailBox:      cf.MailBoxName,
		IsConsumable: cf.IsConsummable,
		Balance:      make(map[common.Address]*StampRecord),
	}, nil
}

type StampConf struct {
	Addr         string
	MailBox      string
	IsConsumable bool
	Balance      map[common.Address]*StampRecord
}

func (sc *StampConf) String() string {
	bts, _ := json.Marshal(sc)
	return string(bts)
}
func (sc *StampConf) GetBalance(ethAddr common.Address, force bool) *StampRecord {

	sr, ok := sc.Balance[ethAddr]
	if ok && !force {
		return sr
	}
	data, err := StampBalanceOfWallet(common.HexToAddress(sc.Addr), ethAddr)
	if err != nil {
		_scLog.Warn(err)
		return nil
	}
	sc.Balance[ethAddr] = &data
	return &data
}
