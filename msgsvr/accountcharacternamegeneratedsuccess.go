package msgsvr

import (
	"fmt"

	"github.com/kralamoure/d1proto"
)

type AccountCharacterNameGeneratedSuccess struct {
	Name string
}

func (m AccountCharacterNameGeneratedSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterNameGeneratedSuccess
}

func (m AccountCharacterNameGeneratedSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Name), nil
}

func (m *AccountCharacterNameGeneratedSuccess) Deserialize(extra string) error {
	m.Name = extra

	return nil
}
