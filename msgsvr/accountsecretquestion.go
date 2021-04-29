package msgsvr

import (
	"net/url"

	"github.com/kralamoure/d1encoding"
)

type AccountSecretQuestion struct {
	Value string
}

func (m AccountSecretQuestion) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountSecretQuestion
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
