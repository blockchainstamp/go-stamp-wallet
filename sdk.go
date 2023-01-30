package bstamp

import (
	"github.com/blockchainstamp/go-stamp-wallet/comm"
	"github.com/blockchainstamp/go-stamp-wallet/sdk_v1"
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

type StampSdk interface {
	CreateWallet(auth string) (comm.Wallet, error)
	GetWallet(addr comm.Address) (comm.Wallet, error)
	ListAllWalletAddr() string
	PrepareWallet(walletAddr comm.Address, auth string) error
	CreateStamp(from, mID string) (comm.Stamp, error)
	ActiveStamp(user, stamp string) error
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
