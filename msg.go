package wechat

import (
	"encoding/xml"
	"fmt"
	"time"
)

type MsgInfo struct {
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
	msg.Load(data)
	if arc, ok := msg.(Archiver); ok {
		err := arc.Archive()
		if err != nil {
			log.Errorf("Archived err, ", err)
		}
	}
	return msg.MsgHandle()

}

func (m *MsgInfo) getResource() (msg MsgHandler) {
	switch m.MsgType {
	case "text":
		msg = new(TextMsg)
	case "image":
		msg = new(ImageMsg)
	case "voice":
		msg = new(VoiceMsg)
	case "video":
		msg = new(VideoMsg)
	case "location":
		msg = new(LocationMsg)
	case "link":
		msg = new(LinkMsg)
	case "event":
		switch m.Event {
		case "subscribe":
			msg = new(SubscribeEvent)
		case "unsubscribe":
			msg = new(UnsubscribeEvent)
		case "SCAN":
			msg = new(ScanEvent)
		case "LOCATION":
			msg = new(LocationEvent)
		case "CLICK":
			msg = new(ClickEvent)
		case "VIEW":
			msg = new(ViewEvent)
		}
	default:
		return nil
	}
	return
}

func (m *MsgInfo) WriteText(data string) {
	xmlfmt := `<xml><ToUserName><![CDATA[%s]]>CD</ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>`
	body := fmt.Sprintf(xmlfmt, m.FromUserName, m.ToUserName, time.Now().Unix(), data)
	log.Info("receive body: ", string(m.content))
	log.Info("send body: ", body)
	m.w.WriteHeader(http.StatusOK)
	m.w.Write([]byte(body))
}
