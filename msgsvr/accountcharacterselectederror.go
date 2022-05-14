package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterSelectedError struct{}

func (m AccountCharacterSelectedError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterSelectedError
}

func (m AccountCharacterSelectedError) MessageName() string {
	return "AccountCharacterSelectedError"
}

func (m AccountCharacterSelectedError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterSelectedError) Deserialize(extra string) error {
	return nil
}
