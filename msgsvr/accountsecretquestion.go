package msgsvr

import (
	"fmt"
	"net/url"

	"github.com/kralamoure/retroproto"
)

type AccountSecretQuestion struct {
	Value string
}

func NewAccountSecretQuestion(extra string) (AccountSecretQuestion, error) {
	var m AccountSecretQuestion

	if err := m.Deserialize(extra); err != nil {
		return AccountSecretQuestion{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSecretQuestion) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountSecretQuestion
}

func (m AccountSecretQuestion) MessageName() string {
	return "AccountSecretQuestion"
}

func (m AccountSecretQuestion) Serialized() (string, error) {
	return url.QueryEscape(m.Value), nil
}

func (m *AccountSecretQuestion) Deserialize(extra string) error {
	v, err := url.QueryUnescape(extra)
	if err != nil {
		return err
	}
	m.Value = v
	return nil
}
