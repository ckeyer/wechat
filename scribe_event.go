package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type ScribeEvent struct {
	Msg
}

func (m *Msg) ReceiveSubscribeEvent() string {
	var msg SubscribeEvent
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (s *ScribeEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(s)
	if err == nil {
		s.Id = id
	}
	return err
}
