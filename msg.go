package wechat

import (
	"encoding/xml"
	"fmt"
)

type MsgInfo struct {
	ID           int64  `xml:"omitempty",orm:"id"`
	MsgType      string `xml:"MsgType,cdata",orm:"MsgType"`
	Event        string `xml:"Event",orm:"Event"`
	ToUserName   string `xml:"ToUserName,cdata",orm:"ToUserName"`
	FromUserName string `xml:"FromUserName,cdata",orm:"FromUserName"`
	CreateTime   int    `xml:"CreateTime",orm:"CreateTime"`
}

func MsgHandle(data []byte) (interface{}, error) {
	v := &MsgInfo{}
	var bs []byte
	copy(bs, data)
	err := xml.Unmarshal(bs, v)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	msg := v.getResource()
	if msg == nil {
		err := fmt.Errorf("%s", "Unknown message type.")
		log.Error(err)
		return nil, err
	}

	if err := xml.Unmarshal(data, msg); err != nil {
		return nil, err
	}

	if arc, ok := msg.(Archiver); ok {
		err := arc.Archive()
		if err != nil {
			log.Errorf("Archived err, ", err)
		}
	}
	return msg.MsgHandle()
}

func (m *MsgInfo) getResource() MsgHandler {
	switch m.MsgType {
	case "text":
		return new(TextMsg)
	case "image":
		return new(ImageMsg)
	case "voice":
		return new(VoiceMsg)
	case "video":
		return new(VideoMsg)
	case "location":
		return new(LocationMsg)
	case "link":
		return new(LinkMsg)
	case "event":
		switch m.Event {
		case "subscribe":
			return new(SubscribeEvent)
		case "unsubscribe":
			// TODO; not sure
			return new(ScribeEvent)
		case "SCAN":
			return new(ScanEvent)
		case "LOCATION":
			return new(LocationEvent)
		case "CLICK":
			return new(ClickEvent)
		case "VIEW":
			return new(ViewEvent)
		}
	}
	return nil
}

// default auto response
func (m *MsgInfo) MsgHandle() (interface{}, error) {
	return NewTextResposeMessage(m.ToUserName, m.FromUserName, "Go Go Go!!!"), nil
}
