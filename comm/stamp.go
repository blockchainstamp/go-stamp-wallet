package comm

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

type Stamp interface {
	Serial() string
	RawData() StampData
	SigData() StampSig
}
type StampData interface {
	SetMsgID(id string)
	GetMailSender() string
	SetNo(no int)
	GetWalletAddr() WalletAddr
	GetStampAddr() StampAddr
	GetMailID() string
	SetEthAddr(addr common.Address)
	SetAddr(addr WalletAddr)
	IsValidInitData() bool
}
type StampSig interface {
	Data() string
	Suffix() string
}

type SimpleStampSig struct {
	SigData   string
	PubSuffix string
}

func (sss *SimpleStampSig) Data() string {
	return sss.SigData
}
func (sss *SimpleStampSig) Suffix() string {
	return sss.PubSuffix
}

type SimpleStamp struct {
	Data *RawStamp
	Sig  *SimpleStampSig
}

func (ss *SimpleStamp) RawData() StampData {
	return ss.Data
}

func (ss *SimpleStamp) SigData() StampSig {
	return ss.Sig
}

func (ss *SimpleStamp) Serial() string {
	bts, _ := json.Marshal(ss)
	return string(bts)
}

type RawStamp struct {
	WAddr        WalletAddr     `json:"wallet_addr"`
	SAddr        StampAddr      `json:"stamp_addr"`
	EAddr        common.Address `json:"eth_addr"`
	FromMailAddr string         `json:"from_mail_addr"`
	MsgID        string         `json:"msg_id"`
	No           int            `json:"no"`
	Time         int64          `json:"time"`
}

func (r *RawStamp) GetWalletAddr() WalletAddr {
	return r.WAddr
}

func (r *RawStamp) GetStampAddr() StampAddr {
	return r.SAddr
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

func (r *RawStamp) GetMailID() string {
	return r.MsgID
}
func (r *RawStamp) GetMailSender() string {
	return r.FromMailAddr
}

func (r *RawStamp) IsValidInitData() bool {
	return len(r.SAddr) > 0 && len(r.FromMailAddr) > 0 && len(r.MsgID) > 0
}
func (r *RawStamp) SetAddr(addr WalletAddr) {
	r.WAddr = addr
}
