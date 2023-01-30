package sdk_v1

import (
	"errors"
	"github.com/blockchainstamp/go-stamp-wallet/comm"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	AllWalletDBKey     = "_stamp_all_wallet_db_key_"
	WalletKeySeparator = '|'
)

var (
	SErrNoSuchWallet = errors.New("no such wallet")
)

type SDK struct {
	database *comm.SDataBase
	wallets  map[comm.Address]comm.Wallet
}

func NewSdk(db *leveldb.DB) *SDK {
	sdk := &SDK{
		database: &comm.SDataBase{DB: db},
		wallets:  make(map[comm.Address]comm.Wallet),
	}
	return sdk
}

func (sdk *SDK) PrepareWallet(walletAddr comm.Address, auth string) error {
	var wallet = &comm.SimpleWallet{}
	err := sdk.database.GetJsonObj([]byte(walletAddr), wallet)
	if err != nil {
		return err
	}
	err = wallet.Open(auth)
	if err != nil {
		return err
	}
	sdk.wallets[walletAddr] = wallet
	return nil
}

func (sdk *SDK) CreateWallet(auth string) (comm.Wallet, error) {
	w, e := comm.CreateWallet(auth)
	if e != nil {
		return nil, e
	}
	key := []byte(w.Address())
	e = sdk.database.SaveJsonObj(key, w)
	if e != nil {
		return nil, e
	}

	e = sdk.database.AppendString(AllWalletDBKey, WalletKeySeparator, string(w.Address()))
	if e != nil {
		return nil, e
	}
	return w, nil
}

func (sdk *SDK) CreateStamp(from, mID string) (comm.Stamp, error) {
	s := comm.NewStamp(from, mID)
	//TODO::
	return s, nil
}

func (sdk *SDK) GetWallet(addr comm.Address) (comm.Wallet, error) {
	var wallet = &comm.SimpleWallet{}
	err := sdk.database.GetJsonObj([]byte(addr), wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (sdk *SDK) ActiveStamp(user, stamp string) error {
	return nil
}
func (sdk *SDK) ListAllWalletAddr() string {
	return sdk.database.GetString(AllWalletDBKey)
}
