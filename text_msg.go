package wechat

type TextMsg struct {
	MsgInfo
	Content string `xml:"Content"`
	MsgId   int64  `xml:"MsgId"`
}
