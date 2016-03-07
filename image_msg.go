package wechat

import (
	"encoding/xml"
	"github.com/astaxie/beego/orm"
)

type ImageMsg struct {
	MsgInfo
	PicUrl  string `xml:"PicUrl"`
	MediaId int    `xml:"MediaId"`
	MsgId   int64  `xml:"MsgId"`
}

func (i *ImageMsg) Insert() error {
	o := orm.NewOrm()
	_, e := o.Insert(i)
	if e != nil {
		return e
	}
	return nil
}

func (m *Msg) ReceiveImageMsg() string {
	var msg ImageMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return ""
	}
	return ""
}
