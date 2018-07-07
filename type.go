package miniprogram

const (
	JsCode2SessionUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	GetAccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type User struct {
	OpenId   string `json:"open_id"`
	NickName string `json:"nick_name"`
	HeadUrl  string `json:"head_url"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Sex      int    `json:"gender"`
	Language string `json:"language"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
	Unionid  string `json:"-"`
}

type GetUserArgs struct {
	Code          string `json:"code"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
}

type Session struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ExpiresIn  int    `json:"expires_in"`
}
