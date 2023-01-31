package comm

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

const (
	WalletVersion = "1"
)

var (
	WOpenErr = errors.New("open wallet failed")
)

type Wallet interface {
	Address() WalletAddr
	Open(auth string) error
	Verbose() string
	Close()
	Sign(s StampData) StampSig
	IsOpen() bool
	EthAddr() common.Address
}

type SimpleWallet struct {
	Version string         `json:"version"`
	Addr    WalletAddr     `json:"address"`
	EAddr   common.Address `json:"eth_addr"`
	Cipher  *CipherData    `json:"cipher"`
	priKey  ed25519.PrivateKey
}

func CreateWallet(auth string) (Wallet, error) {
	pub, pri, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	cipherTxt, err := encryptPriKey(pri, auth)
	if err != nil {
		return nil, err
	}
	ethPri := crypto.ToECDSAUnsafe(pri[:crypto.DigestLength])

	sw := &SimpleWallet{
		Version: WalletVersion,
		Cipher:  cipherTxt,
		Addr:    PubToAddr(pub),
		EAddr:   crypto.PubkeyToAddress(ethPri.PublicKey),
		priKey:  pri,
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
	addr := PubToAddr(pub)
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
	return nil
}
func (sw *SimpleWallet) IsOpen() bool {
	return sw.priKey != nil
}
func (sw *SimpleWallet) EthAddr() common.Address {
	return sw.EAddr
}
