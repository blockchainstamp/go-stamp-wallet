package comm

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

const (
	WalletVersion = "1"
)

var (
	WOpenErr    = errors.New("open wallet failed")
	WVerifyErr  = errors.New("verify signature failed")
	WInvalidSig = errors.New("invalid signature data")
)

type Wallet interface {
	Name() string
	SetName(nn string)
	Address() WalletAddr
	Open(auth string) error
	Verbose() string
	Close()
	Sign(s StampData) StampSig
	IsOpen() bool
	EthAddr() common.Address
	fmt.Stringer
}

type SimpleWallet struct {
	NickName string         `json:"nick_name"`
	Version  string         `json:"version"`
	Addr     WalletAddr     `json:"address"`
	EAddr    common.Address `json:"eth_addr"`
	Cipher   *CipherData    `json:"cipher"`
	priKey   ed25519.PrivateKey
}

func CreateWallet(auth, name string) (Wallet, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	cipherTxt, err := encryptPriKey(pri, auth)
	if err != nil {
		return nil, err
	}
	ethPri := crypto.ToECDSAUnsafe(pri[:crypto.DigestLength])
	addr, _ := PubToAddr(pub)
	sw := &SimpleWallet{
		NickName: name,
		Version:  WalletVersion,
		Cipher:   cipherTxt,
		Addr:     addr,
		EAddr:    crypto.PubkeyToAddress(ethPri.PublicKey),
		priKey:   pri,
	}

	return sw, nil
}

func LoadByJsonData(jsonStr string) (Wallet, error) {
	w := new(SimpleWallet)
	if err := json.Unmarshal([]byte(jsonStr), w); err != nil {
		return nil, err
	}
	return w, nil
}

func LoadByFile(path string) (Wallet, error) {
	bts, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	w := new(SimpleWallet)
	if err := json.Unmarshal(bts, w); err != nil {
		return nil, err
	}
	return w, nil
}

func (sw *SimpleWallet) Close() {
	sw.priKey = nil
}

func (sw *SimpleWallet) Open(auth string) error {
	if sw.priKey != nil {
		return nil
	}

	pri, err := decryptPriKey(sw.Cipher, auth)
	if err != nil {
		return err
	}
	pub := pri.Public().(ed25519.PublicKey)
	addr, _ := PubToAddr(pub)
	if addr != sw.Addr {
		return WOpenErr
	}
	sw.priKey = pri
	return nil
}

func (sw *SimpleWallet) Address() WalletAddr {
	return sw.Addr
}
func (sw *SimpleWallet) Verbose() string {
	bts, _ := json.MarshalIndent(sw, "", "\t")
	return string(bts)
}

func (sw *SimpleWallet) Sign(s StampData) StampSig {
	rawBytes, _ := json.Marshal(s)
	pub := sw.priKey.Public().(ed25519.PublicKey)
	sig := ed25519.Sign(sw.priKey, rawBytes)
	_, suffix := PubToAddr(pub)
	return &SimpleStampSig{
		SigData:   hex.EncodeToString(sig),
		PubSuffix: suffix,
	}
}

func VerifyStamp(stamp Stamp) error {
	var (
		err               error
		data, sig, pubBts []byte
	)
	if stamp.RawData() == nil || stamp.SigData() == nil {
		return WInvalidSig
	}
	data, _ = json.Marshal(stamp.RawData())

	sig, err = hex.DecodeString(stamp.SigData().Data())
	if err != nil {
		return err
	}

	pubBts, err = RecoverPub(stamp.RawData().GetWalletAddr(), stamp.SigData().Suffix())
	if err != nil {
		return err
	}
	if len(pubBts) != ed25519.PublicKeySize {
		return WVerifyErr
	}
	if !ed25519.Verify(pubBts, data, sig) {
		return WVerifyErr
	}
	return nil
}

func (sw *SimpleWallet) IsOpen() bool {
	return sw.priKey != nil
}
func (sw *SimpleWallet) EthAddr() common.Address {
	return sw.EAddr
}
func (sw *SimpleWallet) String() string {
	bs, _ := json.MarshalIndent(sw, "", "\t")
	return string(bs)
}
func (sw *SimpleWallet) Name() string {
	return sw.NickName
}
func (sw *SimpleWallet) SetName(newName string) {
	sw.NickName = newName
}
