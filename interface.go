package wechat

type MsgHandler interface {
	MsgHandle() (interface{}, error)
}

type Archiver interface {
	Archive() error
}
