package wechat

type VideoMsg struct {
	MsgInfo

	MediaId      int    `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	MsgId        int64  `xml:"MsgId"`
}
