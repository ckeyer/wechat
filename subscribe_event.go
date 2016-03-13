package wechat

type SubscribeEvent struct {
	MsgInfo
	EventKey int32  `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}
