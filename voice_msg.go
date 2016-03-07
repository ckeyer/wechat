package wechat

import (
	"encoding/xml"
)

type VoiceMsg struct {
	Msg
	MediaId     string `xml:"MediaId"`
	Format      string `xml:"Format"`
	Recognition string `xml:"Recognition"`
	MsgId       int64  `xml:"MsgId"`
}

func (m *Msg) ReceiveVoiceMsg() {
	var msg VoiceMsg
	err := xml.Unmarshal(m.content, &msg)
	if err != nil {
		return
	}
	ret := replier.Reply(msg.Recognition)

	log.Debug("voice format: ", msg.Format)
	log.Debug("voicd recognition: ", msg.Recognition)
	log.Notice("replied: ", ret)
	m.WriteText(ret)

	m.WriteText("语音识别为：" + msg.Recognition)
	return
}

func (v *VoiceMsg) Insert() error {
	return nil
}
