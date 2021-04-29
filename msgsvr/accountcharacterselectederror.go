package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterSelectedError struct{}

func (m AccountCharacterSelectedError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterSelectedError
}

func (m AccountCharacterSelectedError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterSelectedError) Deserialize(extra string) error {
	return nil
}
