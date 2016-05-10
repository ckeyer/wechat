package wechat

import (
	"encoding/xml"
	"fmt"
	"time"
)

type MsgInfo struct {
	ID           int64 `xml:"omitempty",orm:"id"`
	MsgType      CData `xml:"MsgType",orm:"MsgType"`
	Event        CData `xml:"Event",orm:"Event"`
	ToUserName   CData `xml:"ToUserName",orm:"ToUserName"`
	FromUserName CData `xml:"FromUserName",orm:"FromUserName"`
	CreateTime   int   `xml:"CreateTime",orm:"CreateTime"`
}

func MsgHandle(data []byte) (rm *ResponseMessage, err error) {
	v := &MsgInfo{}
	bs := data[:]

	err = xml.Unmarshal(bs, v)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("msginfo: %#v", v)

	msg := v.getResource()
	if msg == nil {
		err = fmt.Errorf("%s", "Unknown message type.")
		log.Error(err)
		return
	}

	if err = xml.Unmarshal(data, msg); err != nil {
		return
	}
	log.Debugf("msg: %#v", msg)

	if arc, ok := msg.(Archiver); ok {
		err = arc.Archive()
		if err != nil {
			log.Errorf("Archived err, ", err)
		}
	}
	return msg.MsgHandle()
}

func (m *MsgInfo) getResource() MsgHandler {
	switch m.MsgType.Content {
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
		switch m.Event.Content {
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
func (m *MsgInfo) MsgHandle() (*ResponseMessage, error) {
	return NewTextResposeMessage(m.ToUserName.Content, m.FromUserName.Content, time.Now().String()), nil
}
