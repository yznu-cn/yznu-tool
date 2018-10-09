package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yznu-cn/yznu-tool/yznu-go/config"
	"github.com/yznu-cn/yznu-tool/yznu-go/models"
)

type Code2sessionResult struct {
	ErrorCode  int    `json:"errcode"`
	ErrorMsg   string `json:"errmsg,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	ExpiresIn  int    `json:"expires_in,omitempty"`
	Openid     string `json:"openid,omitempty"`
}

func Login(code string, userInfo *models.UserInfo) (token string, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + config.WxConf.AppId + "&secret=" + config.WxConf.Secret + "&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	var m *Code2sessionResult
	r, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(r, &m); err != nil {
		fmt.Println("unmal err : ", err.Error())
		return
	}
	token, err = GenToken("0", userInfo.NickName, userInfo.Openid)
	return
}
