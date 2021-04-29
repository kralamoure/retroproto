package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountCharacterAddSuccess struct{}

func (m AccountCharacterAddSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterAddSuccess
}

func (m AccountCharacterAddSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountCharacterAddSuccess) Deserialize(extra string) error {
	return nil
}
