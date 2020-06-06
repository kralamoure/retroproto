package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterDeleteError struct{}

func (m AccountCharacterDeleteError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterDeleteError
}

func (m AccountCharacterDeleteError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterDeleteError) Deserialize(extra string) error {
	return nil
}
