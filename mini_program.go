package miniprogram

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type MiniProgramInterface interface {
	GetUser(args *GetUserArgs) (*User, error)
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

func (m *MiniProgramImpl) GetUser(args *GetUserArgs) (*User, error) {
	session, err := m.GetSession(args.Code)
	if err != nil {
		return nil, err
	}
	user := &User{}
	wxBizDataCrypt := WxBizDataCrypt{m.AppId, session.SessionKey}
	j, err := wxBizDataCrypt.Decrypt(args.EncryptedData, args.Iv, true)
	if err != nil {
		return nil, err
	}
	s, _ := j.(string)
	err = json.Unmarshal([]byte(s), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
