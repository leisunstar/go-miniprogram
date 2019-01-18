package miniprogram

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type MiniProgramInterface interface {
	GetSession(code string) (*Session, error)
	Decode(encryptedData, iv string, v interface{}) error
	GetWXacodeunLimit(scene, page string, width int, isHyaline bool, filePath string) error
	GetWXacodeunLimitWriter(scene, page string, width int, isHyaline bool, writer io.Writer)
	GetAccessToken() (string, error)
}

type MiniProgramImpl struct {
	mu          sync.Mutex
	AccessToken string
	Expires     int64
	AppId       string
	AppSecret   string
}

func NewMiniProgramImpl() *MiniProgramImpl {
	return &MiniProgramImpl{}
}

func (m *MiniProgramImpl) AddAppId(appId string) *MiniProgramImpl {
	m.AppId = appId
	return m
}

func (m *MiniProgramImpl) AddAppSecret(secret string) *MiniProgramImpl {
	m.AppSecret = secret
	return m
}

func (m *MiniProgramImpl) GetSession(code string) (*Session, error) {
	s := &Session{}
	resp, body, errs := gorequest.New().Get(fmt.Sprintf(JsCode2SessionUrl,
		m.AppId, m.AppSecret, code)).EndStruct(s)
	if errs != nil {
		return nil, fmt.Errorf("get session error %v", errs)
	}
	if len(s.OpenId) == 0 {
		return nil, fmt.Errorf("get session error %s body %s", resp.Status, body)
	}
	return s, nil
}

func (m *MiniProgramImpl) Decode(encryptedData, iv string, session *Session, v interface{}) error {
	wxBizDataCrypt := WxBizDataCrypt{m.AppId, session.SessionKey}
	j, err := wxBizDataCrypt.Decrypt(encryptedData, iv, true)
	if err != nil {
		return err
	}
	s, _ := j.(string)
	err = json.Unmarshal([]byte(s), v)
	if err != nil {
		return err
	}
	return nil
}

func (m *MiniProgramImpl) GetWXacodeunLimitToFile(scene, page string, width int, isHyaline bool, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return m.GetWXacodeunLimitWriter(scene, page, width, isHyaline, f)
}

func (m *MiniProgramImpl) GetWXacodeunLimitWriter(scene, page string, width int, isHyaline bool, writer io.Writer) error {
	token, err := m.GetAccessToken()
	if err != nil {
		return err
	}
	s := &QrCode{
		Scene:     scene,
		Page:      page,
		Width:     width,
		IsHyaline: isHyaline,
	}
	var errs []error
	resp, body, errs := gorequest.New().Post(fmt.Sprintf(GetWxacodeunLimitUrl, token)).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		SendStruct(s).
		End()
	if errs != nil {
		return fmt.Errorf("get qrcode error %v", errs)
	}
	if len(body) < 1000 {
		return fmt.Errorf("get qrcode error %s body %s", resp.Status, body)
	}
	_, err = writer.Write([]byte(body))
	return err
}

func (m *MiniProgramImpl) GetAccessToken() (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.AccessToken) > 0 {
		if time.Now().Unix() < m.Expires {
			return m.AccessToken, nil
		}
	}
	ac := &AccessToken{}
	resp, body, errs := gorequest.New().Get(fmt.Sprintf(GetAccessTokenUrl,
		m.AppId, m.AppSecret)).EndStruct(ac)
	if errs != nil {
		return "", fmt.Errorf("%v", errs)
	}
	if len(ac.Token) == 0 {
		return "", fmt.Errorf("getAccessToken error %ds body %s", resp.Status, body)
	}
	m.AccessToken = ac.Token
	m.Expires = time.Now().Unix() + int64(ac.ExpiresIn)
	return m.AccessToken, nil
}
