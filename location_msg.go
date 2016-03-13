package wechat

type LocationMsg struct {
	MsgInfo
	Location_X float64 `xml: "Location_X"`
	Location_Y float64 `xml: "Location_Y"`
	Scale      int     `xml: "Scale"`
	Label      string  `xml: "Label"`
	MsgId      int64   `xml: "MsgId"`
}
