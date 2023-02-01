package contract

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

var (
	addr   = ""
	user   = ""
	priKey = ""
	val    int64
)

func init() {
	flag.StringVar(&addr, "addr", "0x20B10D21d54D761238BFCD372c2A146586145343", "--addr")
	flag.StringVar(&user, "uid", "", "--uid")
	flag.StringVar(&priKey, "pri", "", "--pri")
	flag.Int64Var(&val, "val", 0, "--val")
}
func TestStampBalanceOfWallet(t *testing.T) {
	fmt.Println(StampBalanceOfWallet(common.HexToAddress(addr), common.HexToAddress(user)))
}

func TestTransferToUser(t *testing.T) {
	bts, err := hex.DecodeString(priKey)
	if err != nil {
		panic(err)
	}
	pri := crypto.ToECDSAUnsafe(bts)
	caller, op, err := EthTxCli(addr, pri)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := caller.Transfer(op, common.HexToAddress(user), big.NewInt(val))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Sprintln(tx.Hash().String())
}
