package go_stamp_wallet

type StampAddress string

type Wallet interface {
	Addr() StampAddress
	Balance() int64
	Open(auth string)
}
type Stamp interface {
	MailAddr() string
	FromWallet() string
	ToWallet() string
	MailID() string
	SignCli() string
	SignSrv() string
}
