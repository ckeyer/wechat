package wechat

type ViewEvent struct {
	MsgInfo
	EventKey string `xml:"EventKey"`
}
