/*
 * 与腾讯服务器的相关交互
**/
package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	//  redis中微信 AccessToken 的key
	REDIS_KEY_WC_ACCESS_TOKEN = "wx_AccessToken"
)

type AccessToken struct {
	AccessToken string `json: "access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	created     time.Time
}

// 更新微信的AccessToken到Redis中 key=REDIS_KEY_WC_ACCESS_TOKEN
func UpdateAccessToken() (expires_in int, err error) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" +
		config.WeChat.AppId + "&secret=" + config.WeChat.AppSecret

	if req, err := http.Get(url); err != nil {
		log.Error(err)
		return 0, err
	} else {
		var v AccessToken
		bs, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Error(err)
			return 0, err
		}

		err = json.Unmarshal(bs, &v)
		if err != nil {
			log.Error(err)
			return 0, err
		}
		v.created = time.Now()

		access_token = v.AccessToken
		expires_in = (int)(v.ExpiresIn)

		log.Notice("Successful: get AccessToken ")
	}
	return
}

func AutoGetAccessToken() {
	ei, err := UpdateAccessToken()
	if err != nil {
		log.Warning(err.Error())
		return
	}
	outtime := (time.Duration)(ei-100) * time.Second
	go time.AfterFunc(outtime, AutoGetAccessToken)
}

func GetAccessToken() string {
	return access_token
}
