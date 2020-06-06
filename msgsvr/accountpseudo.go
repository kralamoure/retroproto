package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountPseudo struct {
	Value string
}

func (m AccountPseudo) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountPseudo
}

func (m AccountPseudo) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AccountPseudo) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
