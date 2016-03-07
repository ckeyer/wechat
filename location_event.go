package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

// 地理位置推送事件结构体
type LocationEvent struct {
	MsgInfo
	Latitude  float64 `xml:"Latitude"`
	Longitude float64 `xml:"Longitude"`
	Precision int     `xml:"Precision"`
}

func (m *Msg) ReceiveLocationEvent() {
	var msg LocationEvent
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return
	}
}

func (l *LocationEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(l)
	if err == nil {
		l.Id = id
	}
	return err
}
