package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type VideoMsg struct {
	Msg
	MediaId      int    `xml:"MediaId"`
	ThumbMediaId string `xml:"ThumbMediaId"`
	MsgId        int64  `xml:"MsgId"`
}

func (m *Msg) ReceiveVideoMsg() string {
	var msg VideoMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}

func (v *VideoMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(v)
	if e != nil {
		return e
	}
	return nil
}
