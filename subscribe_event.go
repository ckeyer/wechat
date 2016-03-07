package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type SubscribeEvent struct {
	Msg
	EventKey int32  `xml:"EventKey"`
	Ticket   string `xml:"Ticket"`
}

func (m *Msg) ReceiveUnsubscribeEvent() string {
	var msg ScribeEvent
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (s *SubscribeEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(s)
	if err == nil {
		s.Id = id
	}
	return err
}
