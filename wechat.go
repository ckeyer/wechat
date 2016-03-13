package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	"github.com/astaxie/beego/orm"
	"github.com/ckeyer/alone/conf"
	logpkg "github.com/ckeyer/go-log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	config       *conf.Config
	log          = logpkg.GetDefaultLogger("wechat")
	access_token string
)

func init() {
	config = conf.GetConfig()
	if config == nil {
		log.Panic("config is nil")
	}
	//	RegDB()
}

func RegDB() {
	// orm.RegisterDataBase("default", "mysql", config.Mysql.GetConnStr())

	orm.RegisterModel(new(TextMsg),
		new(ImageMsg),
		new(LinkMsg),
		new(LocationMsg),
		new(VideoMsg),
		new(VoiceMsg))
}

// Auth 服务器验证
func Auth(signature, timestamp, nonce, echostr string) bool {
	tmps := []string{config.WeChat.Token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]

	tmp := func(data string) string {
		t := sha1.New()
		io.WriteString(t, data)
		return fmt.Sprintf("%x", t.Sum(nil))
	}(tmpStr)

	if tmp == signature {
		return true
	}
	log.Debug("auth receive: ", tmp)
	log.Debug("auth should be: ", signature)
	return false
}
