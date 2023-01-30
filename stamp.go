package go_stamp_wallet

type Stamp interface {
	MailAddr() string
	FromWallet() string
	ToWallet() string
	MailID() string
	SignCli() string
	SignSrv() string
}
