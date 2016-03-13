package wechat

type VoiceMsg struct {
	MsgInfo

	MediaId     string `xml:"MediaId"`
	Format      string `xml:"Format"`
	Recognition string `xml:"Recognition"`
	MsgId       int64  `xml:"MsgId"`
}
