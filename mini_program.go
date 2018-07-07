package miniprogram

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type MiniProgramInterface interface {
	GetSession(code string) (*Session, error)
	Decode(encryptedData, iv string, v interface{}) error
}

type MiniProgramImpl struct {
	AppId     string
	AppSecret string
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
	_, _, errs := gorequest.New().Get(fmt.Sprintf(JsCode2SessionUrl,
		m.AppId, m.AppSecret, code)).EndStruct(s)
	if errs != nil {
		return nil, errors.New(fmt.Sprintf("%v", errs))
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
