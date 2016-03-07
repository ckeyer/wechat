package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type LocationMsg struct {
	MsgInfo
	Location_X float64 `xml: "Location_X"`
	Location_Y float64 `xml: "Location_Y"`
	Scale      int     `xml: "Scale"`
	Label      string  `xml: "Label"`
	MsgId      int64   `xml: "MsgId"`
}

func (m *Msg) ReceiveLocationMsg() string {
	var msg LocationMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (l *LocationMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(l)
	if e != nil {
		return e
	}
	return nil
}
