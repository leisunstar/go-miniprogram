package miniprogram

const (
	JsCode2SessionUrl    = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	GetAccessTokenUrl    = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	GetWxacodeunLimitUrl = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s"
)

type User struct {
	Openid    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	UnionId   string `json:"unionId"`
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
	UnionId    string `json:"unionid"`
}

type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

type QrCode struct {
	Scene     string `json:"scene"`
	Page      string `json:"page"`
	Width     int    `json:"width"`
	IsHyaline bool   `json:"is_hyaline"`
}

type QrCodeArgs struct {
	QrCode
	FilePath string `json:"file_path"`
}
