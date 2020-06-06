package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterSelectedError struct{}

func (m AccountCharacterSelectedError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterSelectedError
}

func (m AccountCharacterSelectedError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterSelectedError) Deserialize(extra string) error {
	return nil
}
