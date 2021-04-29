package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterDeleteError struct{}

func (m AccountCharacterDeleteError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterDeleteError
}

func (m AccountCharacterDeleteError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterDeleteError) Deserialize(extra string) error {
	return nil
}
