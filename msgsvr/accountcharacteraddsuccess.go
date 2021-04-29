package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountCharacterAddSuccess struct{}

func (m AccountCharacterAddSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterAddSuccess
}

func (m AccountCharacterAddSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterAddSuccess) Deserialize(extra string) error {
	return nil
}
