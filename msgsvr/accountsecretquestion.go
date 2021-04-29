package msgsvr

import (
	"net/url"

	"github.com/kralamoure/d1proto"
)

type AccountSecretQuestion struct {
	Value string
}

func (m AccountSecretQuestion) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountSecretQuestion
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
