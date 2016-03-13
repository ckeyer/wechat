package wechat

// 地理位置推送事件结构体
type LocationEvent struct {
	MsgInfo
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision int     `xml:"Precision"`
}
