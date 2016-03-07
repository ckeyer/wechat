package wechat

import (
	"time"
)

type ResponseMessage struct {
	ToUserName   string `xml:"ToUserName,cdata"`
	FromUserName string `xml:"FromUserName,cdata"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType,cdata"`

	Content      string `xml:"Content,cdata,omitempty"`
	MediaId      string `xml:"MediaId,cdata,omitempty"`
	Title        string `xml:"Title,cdata,omitempty"`
	Description  string `xml:"Description,cdata,omitempty"`
	MusicURL     string `xml:"MusicURL,cdata,omitempty"`
	HQMusicUrl   string `xml:"HQMusicUrl,cdata,omitempty"`
	ThumbMediaId string `xml:"ThumbMediaId,cdata,omitempty"`
	ArticleCount int    `xml:"ArticleCount,omitempty"`
	Articles     string `xml:"Articles,cdata,omitempty"`
	PicUrl       string `xml:"PicUrl,cdata,omitempty"`
	Url          string `xml:"Url,cdata,omitempty"`
}

func NewTextResposeMessage(to, from, content string) *ResponseMessage {
	return &ResponseMessage{
		ToUserName:   to,
		FromUserName: from,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}
}
