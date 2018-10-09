package models


type UserInfo struct {
	Openid    string `json:"openid"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	WaterMark `json:"watermark"`
}


type WaterMark struct {
	AppID     string `json:"appid"`
	Timestamp int    `json:"timestamp"`
}