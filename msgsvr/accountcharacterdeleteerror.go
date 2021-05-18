package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterDeleteError struct{}

func (m AccountCharacterDeleteError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterDeleteError
}

func (m AccountCharacterDeleteError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterDeleteError) Deserialize(extra string) error {
	return nil
}
