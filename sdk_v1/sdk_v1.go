package sdk_v1

import (
	"errors"
	"fmt"
	"github.com/blockchainstamp/go-stamp-wallet/comm"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	WalletDetailsDBKEY = "_wallet_details_db_key_"
	StampDetailsDBKEY  = "_stamp_details_db_key_"
	AllWalletDBKey     = "_stamp_all_wallet_db_key_"
	WalletKeySeparator = '|'
)

var (
	SErrNoSuchWallet = errors.New("no such wallet")
	SErrWalletClosed = errors.New("wallet closed")
	SErrInvalid      = errors.New("invalid stamp or wallet")
)

func walletKey(wAddr comm.WalletAddr) []byte {
	s := fmt.Sprintf("%s%s", WalletDetailsDBKEY, wAddr)
	return []byte(s)
}

func stampKey(sAddr comm.StampAddr) []byte {
	s := fmt.Sprintf("%s%s", StampDetailsDBKEY, sAddr)
	return []byte(s)
}

type SDK struct {
	database  *comm.SDataBase
	wallets   map[comm.WalletAddr]comm.Wallet
	stampAddr map[string]comm.StampAddr
}

func NewSdk(db *leveldb.DB) *SDK {
	sdk := &SDK{
		database:  &comm.SDataBase{DB: db},
		wallets:   make(map[comm.WalletAddr]comm.Wallet),
		stampAddr: make(map[string]comm.StampAddr),
	}
	return sdk
}
func (sdk *SDK) PrepareWallet(walletAddr comm.WalletAddr, auth string) error {
	var wallet = &comm.SimpleWallet{}
	err := sdk.database.GetJsonObj(walletKey(walletAddr), wallet)
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
	key := walletKey(w.Address())
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

func (sdk *SDK) SignStamp(sData comm.StampData) (comm.Stamp, error) {
	wAddr := sData.GetWalletAddr()
	sAddr := sData.GetStampAddr()
	if len(wAddr) == 0 || len(sAddr) == 0 {
		return nil, SErrInvalid
	}

	w, ok := sdk.wallets[wAddr]
	if !ok {
		return nil, SErrNoSuchWallet
	}

	if !w.IsOpen() {
		return nil, SErrWalletClosed
	}

	sData.SetEthAddr(w.EthAddr())
	//TODO::
	sig := w.Sign(sData)
	stamp := &comm.SimpleStamp{
		Data: sData,
		Sig:  sig,
	}
	return stamp, nil
}

func (sdk *SDK) GetWallet(addr comm.WalletAddr) (comm.Wallet, error) {
	var wallet = &comm.SimpleWallet{}
	err := sdk.database.GetJsonObj(walletKey(addr), wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (sdk *SDK) ActiveStamp(mailUser string, sAddr comm.StampAddr) error {
	if len(sAddr) == 0 {
		delete(sdk.stampAddr, mailUser)
		return nil
	}
	lastStamp := &comm.SimpleStamp{}
	if err := sdk.database.GetJsonObj(stampKey(sAddr), lastStamp); err != nil {
		if err != leveldb.ErrNotFound {
			return err
		}
	}
	sdk.stampAddr[mailUser] = sAddr
	return nil
}

func (sdk *SDK) GetActiveStamp(user string) comm.StampAddr {
	addr, ok := sdk.stampAddr[user]
	if !ok {
		return ""
	}
	return addr
}

func (sdk *SDK) ListAllWalletAddr() string {
	return sdk.database.GetString(AllWalletDBKey)
}
func (sdk *SDK) UpdateStampBalanceAsync(sAddr comm.StampAddr) {

}
