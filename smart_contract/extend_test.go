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
	flag.StringVar(&addr, "addr", "0xF9Cbfd74808f812a3B8A2337BFC426B2A10Bd74a", "--addr")
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

func TestStampConfigFromBlockChain(t *testing.T) {
	fmt.Println(StampConfigFromBlockChain(addr))
}
