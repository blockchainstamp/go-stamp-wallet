package comm

import (
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"strings"
)

type Address string

const (
	AccPrefix = "BS"
	AccLen    = 20
)

var (
	InvalidAddr = errors.New("invalid address")
)

func PubToAddr(key []byte) Address {
	codedKey := base58.Encode(key)
	if len(codedKey) > AccLen {
		codedKey = string([]byte(codedKey)[:AccLen])
	}
	return Address(AccPrefix + codedKey)
}

func RecoverPub(addr Address, sub string) ([]byte, error) {
	if len(addr) < AccLen {
		return nil, InvalidAddr
	}
	source := string(addr) + sub
	source = strings.TrimPrefix(source, AccPrefix)
	return base58.Decode(source), nil
}
