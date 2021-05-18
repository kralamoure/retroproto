package msgsvr

import (
	"net/url"

	"github.com/kralamoure/retroproto"
)

type AccountSecretQuestion struct {
	Value string
}

func (m AccountSecretQuestion) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountSecretQuestion
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
