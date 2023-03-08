package sdk_v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blockchainstamp/go-stamp-wallet/comm"
	contract "github.com/blockchainstamp/go-stamp-wallet/smart_contract"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	WalletDetailsDBKEY = "_wallet_details_db_key_"
	StampDetailsDBKEY  = "_stamp_details_db_key_"
	AllWalletDBKey     = "_stamp_all_wallet_db_key_V2_"
)

var (
	_sdkLog = logrus.WithFields(logrus.Fields{
		"mode": "stamp sdk",
	})
	SErrActiveWallet = errors.New("no active wallet")
	SErrActiveStamp  = errors.New("no active stamp")
	SErrInvalid      = errors.New("invalid stamp raw data")
	SErrInsufficient = errors.New("insufficient fund")
	SErrWInUsed      = errors.New("wallet is in used")
	SErrWDuplicated  = errors.New("wallet duplicated")
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
	database     *comm.SDataBase
	activeWallet comm.Wallet
	stampConf    map[string]*contract.StampConf
}

func NewSdk(db *leveldb.DB) *SDK {
	sdk := &SDK{
		database:  &comm.SDataBase{DB: db},
		stampConf: make(map[string]*contract.StampConf),
	}
	return sdk
}
func (sdk *SDK) ActiveWallet(walletAddr comm.WalletAddr, auth string) (comm.Wallet, error) {
	var wallet = &comm.SimpleWallet{}
	err := sdk.database.GetJsonObj(walletKey(walletAddr), wallet)
	if err != nil {
		_sdkLog.Warn(err)
		return nil, err
	}
	err = wallet.Open(auth)
	if err != nil {
		_sdkLog.Warn(err)
		return nil, err
	}
	sdk.activeWallet = wallet
	return wallet, nil
}

func (sdk *SDK) CreateWallet(auth, name string) (comm.Wallet, error) {
	w, e := comm.CreateWallet(auth, name)
	if e != nil {
		_sdkLog.Warn(e)
		return nil, e
	}
	key := walletKey(w.Address())
	e = sdk.database.SaveJsonObj(key, w)
	if e != nil {
		_sdkLog.Warn(e)
		return nil, e
	}

	e = sdk.database.AppendString(AllWalletDBKey, string(w.Address()))
	if e != nil {
		_sdkLog.Warn(e)
		return nil, e
	}
	return w, nil
}
func (sdk *SDK) RemoveWallet(addr comm.WalletAddr) error {

	if sdk.activeWallet != nil && addr == sdk.activeWallet.Address() {
		return SErrWInUsed
	}
	ok, _ := sdk.database.Has(walletKey(addr), nil)
	if !ok {
		return nil
	}
	err := sdk.database.TrimString(AllWalletDBKey, string(addr))
	if err != nil {
		return err
	}
	return sdk.database.Del(walletKey(addr))
}

func (sdk *SDK) ImportWallet(walletJson, auth string) (comm.Wallet, error) {
	var wallet = &comm.SimpleWallet{}
	err := json.Unmarshal([]byte(walletJson), wallet)
	if err != nil {
		return nil, err
	}
	err = wallet.Open(auth)
	if err != nil {
		return nil, err
	}

	key := walletKey(wallet.Addr)
	ok, _ := sdk.database.Has(key, nil)
	if ok {
		return nil, SErrWDuplicated
	}
	err = sdk.database.SaveJsonObj(key, wallet)
	if err != nil {
		_sdkLog.Warn(err)
		return nil, err
	}

	err = sdk.database.AppendString(AllWalletDBKey, string(wallet.Address()))
	if err != nil {
		_sdkLog.Warn(err)
		return nil, err
	}
	return wallet, nil
}

func (sdk *SDK) PostStamp(mailUser string, sData comm.StampData) (comm.Stamp, error) {
	rawData := sData.(*comm.RawStamp)
	if !rawData.IsValidInitData() {
		_sdkLog.Warn(SErrInvalid)
		return nil, SErrInvalid
	}
	w := sdk.activeWallet
	if w == nil || !w.IsOpen() {
		_sdkLog.Warn(SErrActiveWallet)
		return nil, SErrActiveWallet
	}
	sConf, ok := sdk.stampConf[mailUser]
	if !ok {
		_sdkLog.Warn(SErrActiveStamp)
		return nil, SErrActiveStamp
	}
	rawData.WAddr = w.Address()
	rawData.EAddr = w.EthAddr()
	if sConf.IsConsumable {
		//TODO:: post stamp to manager
		_sdkLog.Warn("no support now")
		return nil, fmt.Errorf("no support now")
	}

	sig := w.Sign(rawData).(*comm.SimpleStampSig)
	stamp := &comm.SimpleStamp{
		Data: rawData,
		Sig:  sig,
	}
	return stamp, nil
}

func (sdk *SDK) GetWallet(addr comm.WalletAddr) (comm.Wallet, error) {
	var wallet = &comm.SimpleWallet{}
	err := sdk.database.GetJsonObj(walletKey(addr), wallet)
	if err != nil {
		_sdkLog.Warn(err)
		return nil, err
	}

	return wallet, nil
}

func (sdk *SDK) ConfigStamp(mailUser string, sAddr comm.StampAddr) error {
	if len(sAddr) == 0 {
		delete(sdk.stampConf, mailUser)
		return nil
	}
	if sdk.activeWallet == nil {
		return SErrActiveWallet
	}
	if sdk.stampConf[mailUser] != nil {
		return nil
	}

	conf, err := contract.StampConfigFromBlockChain(string(sAddr))
	if err != nil {
		_sdkLog.Warn(err)
		return err
	}
	record := conf.GetBalance(sdk.activeWallet.EthAddr(), true)
	if record == nil {
		_sdkLog.Warn(SErrActiveStamp)
		return SErrActiveStamp
	}
	if record.Value == nil || record.Value.Int64() <= 0 {
		_sdkLog.Warn(SErrInsufficient)
		return SErrInsufficient
	}
	sdk.stampConf[mailUser] = conf
	return nil
}

func (sdk *SDK) GetStampConf(user string) *contract.StampConf {
	return sdk.stampConf[user]
}

func (sdk *SDK) ListAllWalletAddr() string {
	return sdk.database.GetString(AllWalletDBKey)
}

func (sdk *SDK) UpdateStampBalanceAsync(mailUser string) {
	go func() {
		conf, ok := sdk.stampConf[mailUser]
		if !ok {
			return
		}
		if sdk.activeWallet == nil {
			return
		}
		_ = conf.GetBalance(sdk.activeWallet.EthAddr(), true)
	}()
}

func (sdk *SDK) VerifyStamp(stampStr, mailID string) error {
	stamp := &comm.SimpleStamp{}
	err := json.Unmarshal([]byte(stampStr), stamp)
	if err != nil {
		_sdkLog.Warn(err)
		return err
	}
	//if mailID != stamp.RawData().GetMailID() {
	//	return SErrMailVerify
	//}

	return comm.VerifyStamp(stamp)
}
