package responses

type CallbackAuth struct {
	Ticket       string
	FromUserName string
	CreateTime   float64
	MsgType      string
	Event        string
}
