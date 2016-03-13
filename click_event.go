package wechat

// 点击事件结构体
type ClickEvent struct {
	MsgInfo
	EventKey string `xml:"EventKey"`
}
