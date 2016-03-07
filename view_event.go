package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type ViewEvent struct {
	Msg
	EventKey string `xml:"EventKey"`
}

func (m *Msg) ReceiveViewEvent() string {
	var msg ViewEvent
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (v *ViewEvent) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(v)
	if err == nil {
		v.Id = id
	}
	return err
}
