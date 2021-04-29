package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountPseudo struct {
	Value string
}

func (m AccountPseudo) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountPseudo
}

func (m AccountPseudo) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AccountPseudo) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
