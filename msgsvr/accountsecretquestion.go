package msgsvr

import (
	"html"

	"github.com/kralamoure/d1proto"
)

type AccountSecretQuestion struct {
	Value string
}

func (m AccountSecretQuestion) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountSecretQuestion
}

func (m AccountSecretQuestion) Serialized() (string, error) {
	return html.EscapeString(m.Value), nil
}

func (m *AccountSecretQuestion) Deserialize(extra string) error {
	m.Value = html.UnescapeString(extra)

	return nil
}
