package wechat

import (
	"encoding/xml"
)

type MsgHandler interface {
	MsgHandle() (interface{}, error)
}

type Archiver interface {
	Archive() error
}

func (m *MsgHandler) Load(data []byte) error {
	return xml.Unmarshal(data, m)
}
