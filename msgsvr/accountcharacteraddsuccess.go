package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountCharacterAddSuccess struct{}

func (m AccountCharacterAddSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterAddSuccess
}

func (m AccountCharacterAddSuccess) MessageName() string {
	return "AccountCharacterAddSuccess"
}

func (m AccountCharacterAddSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterAddSuccess) Deserialize(extra string) error {
	return nil
}
