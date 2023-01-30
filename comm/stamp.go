package comm

import "encoding/json"

type Stamp interface {
	Serial() string
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
	From  string
	MsgID string
	Sig   StampSig
}

func (ss *SimpleStamp) Serial() string {
	bts, _ := json.Marshal(ss)
	return string(bts)
}

func NewStamp(from, msg string) Stamp {
	return &SimpleStamp{
		From:  from,
		MsgID: msg,
		Sig: &SimpleStampSig{
			SigData:   make([]byte, 10),
			PubSuffix: "TODO::",
		},
	}
}
