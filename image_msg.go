package wechat

type ImageMsg struct {
	MsgInfo
	PicUrl  string `xml:"PicUrl"`
	MediaId int    `xml:"MediaId"`
	MsgId   int64  `xml:"MsgId"`
}
