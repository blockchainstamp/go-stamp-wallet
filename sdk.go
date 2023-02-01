package bstamp

import (
	"github.com/blockchainstamp/go-stamp-wallet/comm"
	"github.com/blockchainstamp/go-stamp-wallet/sdk_v1"
	contract "github.com/blockchainstamp/go-stamp-wallet/smart_contract"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"path"
	"sync"
)

var (
	systemBaseDir             = "."
	database      *leveldb.DB = nil
)

func InitSDK(baseDir string) error {
	systemBaseDir = baseDir
	dbPath := path.Join(systemBaseDir, "/", "database")
	opts := &opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.SnappyCompression,
		Filter:      filter.NewBloomFilter(10),
	}

	db, err := leveldb.OpenFile(dbPath, opts)
	if err != nil {
		return err
	}
	database = db
	return nil
}

func SetLogLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

type StampSdk interface {
	CreateWallet(auth string) (comm.Wallet, error)
	GetWallet(addr comm.WalletAddr) (comm.Wallet, error)
	ListAllWalletAddr() string
	ActiveWallet(walletAddr comm.WalletAddr, auth string) (comm.Wallet, error)
	PostStamp(sData comm.StampData) (comm.Stamp, error)
	SetStamp(user string, stamp comm.StampAddr) error
	GetStamp(user string) *contract.StampConf
	UpdateStampBalanceAsync(mailUser string)
	VerifyStamp(stamp, mail string) error
}

var (
	instance StampSdk = nil
	once     sync.Once
)

func Inst() StampSdk {
	if database == nil {
		panic("please InitSDK first")
	}
	once.Do(func() {
		instance = sdk_v1.NewSdk(database)
	})
	return instance
}
