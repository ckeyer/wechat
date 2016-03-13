package wechat

// 二维码扫码事件结构体
type ScanEvent struct {
	MsgInfo
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}
