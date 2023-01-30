package comm

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"errors"
	"os"
)

const (
	WalletVersion = "1"
)

var (
	WOpenErr = errors.New("open wallet failed")
)

type Wallet interface {
	Address() Address
	Open(auth string) error
	Verbose() string
	Close()
}

type SimpleWallet struct {
	Version string      `json:"version"`
	Addr    Address     `json:"address"`
	Cipher  *CipherData `json:"cipher"`
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
	sw := &SimpleWallet{
		Version: WalletVersion,
		Cipher:  cipherTxt,
		Addr:    PubToAddr(pub),
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
func (sw *SimpleWallet) Address() Address {
	return sw.Addr
}
func (sw *SimpleWallet) Verbose() string {
	bts, _ := json.MarshalIndent(sw, "", "\t")
	return string(bts)
}
