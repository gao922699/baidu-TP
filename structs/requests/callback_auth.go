package requests

type CallbackAuth struct {
	Nonce        string
	TimeStamp    string
	Encrypt      string
	MsgSignature string
}
