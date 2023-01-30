package go_stamp_wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/scrypt"
	"io"
)

type CipherData struct {
	KeyParam  `json:"param"`
	Code      string `json:"code"`
	KeyCrypto string `json:"key_crypto"`
	Salt      string `json:"salt"`
	PriCipher string `json:"pri_cipher"`
}
type KeyParam struct {
	S int `json:"s"`
	N int `json:"n"`
	R int `json:"r"`
	P int `json:"p"`
	L int `json:"l"`
}

var KP = KeyParam{
	S: 8,
	N: 1 << 15,
	R: 8,
	P: 1,
	L: 32,
}

func Decrypt(key []byte, cipherTxt []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherTxt) < aes.BlockSize {
		return nil, fmt.Errorf("cipher text too short")
	}

	iv := cipherTxt[:aes.BlockSize]
	cipherTxt = cipherTxt[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTxt, cipherTxt)

	return cipherTxt, nil
}

func Encrypt(key []byte, plainTxt []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainTxt))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTxt)

	return cipherText, nil
}
func AESKey(salt []byte, password string) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, KP.N, KP.R, KP.P, KP.L)
}

func encryptPriKey(priKey ed25519.PrivateKey, pubKey ed25519.PublicKey, auth string) (string, error) {
	salt := make([]byte, KP.S)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	aesKey, err := AESKey(salt, auth)
	if err != nil {
		return "", err
	}
	ci, err := Encrypt(aesKey, priKey[:])
	if err != nil {
		return "", err
	}

	ciData := &CipherData{
		KeyParam:  KP,
		Code:      "base58",
		KeyCrypto: "scrypt",
		Salt:      base58.Encode(salt),
		PriCipher: base58.Encode(ci),
	}
	bs, err := json.MarshalIndent(ciData, "", "\t")
	if err != nil {
		return "", err
	}
	return string(bs), nil // ,base58.Encode(ci)
}

func PrivateKeyToCurve25519(curve25519Private *[32]byte, privateKey *[64]byte) {
	h := sha512.New()
	h.Write(privateKey[:32])
	digest := h.Sum(nil)

	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	copy(curve25519Private[:], digest)
}

func decryptPriKey(cpTxt, auth string) (ed25519.PrivateKey, error) {

	ciData := &CipherData{}

	err := json.Unmarshal([]byte(cpTxt), ciData)
	if err != nil {
		return nil, err
	}
	salt := base58.Decode(ciData.Salt)
	aesKey, err := AESKey(salt, auth)
	if err != nil {
		return nil, err
	}
	priBytes := base58.Decode(ciData.PriCipher)
	return Decrypt(aesKey, priBytes)
}
