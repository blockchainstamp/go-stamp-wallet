package comm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

type Stamp interface {
	Serial() string
}
type StampData interface {
	SetMsgID(id string)
	SetNo(no int)
	GetWalletAddr() WalletAddr
	GetStampAddr() StampAddr
	SetEthAddr(addr common.Address)
	InitByBlockChain(addr common.Address)
}
type StampSig interface {
	Sig() []byte
	Suffix() string
}

type SimpleStampSig struct {
	SigData   []byte
	PubSuffix string
}

func (sss *SimpleStampSig) Sig() []byte {
	return sss.SigData
}
func (sss *SimpleStampSig) Suffix() string {
	return sss.PubSuffix
}

type SimpleStamp struct {
	Data StampData
	Sig  StampSig
}

func (ss *SimpleStamp) Serial() string {
	bts, _ := json.Marshal(ss)
	return string(bts)
}

type RawStamp struct {
	WAddr        WalletAddr     `json:"wallet_addr"`
	SAdr         StampAddr      `json:"stamp_addr"`
	EAddr        common.Address `json:"eth_addr"`
	FromMailAddr string         `json:"from_mail_addr"`
	MsgID        string         `json:"msg_id"`
	No           int            `json:"no"`
}

func (r *RawStamp) GetWalletAddr() WalletAddr {
	return r.WAddr
}

func (r *RawStamp) GetStampAddr() StampAddr {
	return r.SAdr
}

func (r *RawStamp) SetMsgID(id string) {
	r.MsgID = id
}

func (r *RawStamp) SetNo(no int) {
	r.No = no
}
func (r *RawStamp) SetEthAddr(addr common.Address) {
	r.EAddr = addr
}
func (r *RawStamp) InitByBlockChain(addr common.Address) {

}
