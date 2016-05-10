package wechat

import (
	"encoding/xml"
	"time"
)

type CData struct {
	Content string `xml:",cdata"`
}

type ResponseMessage struct {
	XMLName xml.Name `xml:"xml"`

	ToUserName   *CData `xml:"ToUserName"`
	FromUserName *CData `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      *CData `xml:"MsgType"`

	Content      *CData `xml:"Content,omitempty"`
	MediaId      *CData `xml:"MediaId,omitempty"`
	Title        *CData `xml:"Title,omitempty"`
	Description  *CData `xml:"Description,omitempty"`
	MusicURL     *CData `xml:"MusicURL,omitempty"`
	HQMusicUrl   *CData `xml:"HQMusicUrl,omitempty"`
	ThumbMediaId *CData `xml:"ThumbMediaId,omitempty"`
	ArticleCount int    `xml:"ArticleCount,omitempty"`
	Articles     *CData `xml:"Articles,omitempty"`
	PicUrl       *CData `xml:"PicUrl,omitempty"`
	Url          *CData `xml:"Url,omitempty"`
}

func NewTextResposeMessage(from, to, content string) *ResponseMessage {
	rs := &ResponseMessage{
		ToUserName:   &CData{to},
		FromUserName: &CData{from},
		CreateTime:   time.Now().Unix(),
		MsgType:      &CData{"text"},
		Content:      &CData{content},
	}
	bs, err := xml.Marshal(rs)
	log.Debugf("err %s", err)
	log.Debug(string(bs))
	return rs
}
