package wechat

import (
	"encoding/xml"
)

type TextMsg struct {
	ID int64 `xml:"omitempty",orm:"id"`

	MsgType      string    `xml:"MsgType,cdata",orm:"MsgType"`
	Event        string    `xml:"Event",orm:"Event"`
	ToUserName   string    `xml:"ToUserName,cdata",orm:"ToUserName"`
	FromUserName string    `xml:"FromUserName,cdata",orm:"FromUserName"`
	CreateTime   int       `xml:"CreateTime",orm:"CreateTime"`
	CreatedLocal time.Time `xml:"omitempty",orm:"auto_now_add;type(datetime)"`

	Content string `xml:"Content"`
	MsgId   int64  `xml:"MsgId"`
}

func (t *TextMsg) Insert() error {
	return nil
}

func (t *TextMsg) MsgHandle() (interface{}, error) {
	return NewTextResposeMessage(t.ToUserName, t.FromUserName, "Go Go Go!!!"), nil
}
