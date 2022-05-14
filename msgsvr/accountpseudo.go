package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountPseudo struct {
	Value string
}

func (m AccountPseudo) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountPseudo
}

func (m AccountPseudo) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AccountPseudo) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
