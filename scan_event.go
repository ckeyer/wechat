package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

// 二维码扫码事件结构体
type ScanEvent struct {
	Msg
	EventKey string `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

func (m *Msg) ReceiveScanEvent() string {
	var se ScanEvent
	err := xml.Unmarshal(m.content, &se)
	if err != nil {
		return ""
	}
	return ""
}

func (s *ScanEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(s)
	if err == nil {
		s.Id = id
	}
	return err
}
