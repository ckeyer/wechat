package wechat

type LinkMsg struct {
	MsgInfo
	Title       string `xml: "Title"`
	Description string `xml: "Description"`
	Url         string `xml: "Url"`
	MsgId       int64  `xml: "MsgId"`
}
