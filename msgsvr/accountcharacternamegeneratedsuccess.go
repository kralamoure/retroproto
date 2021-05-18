package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterNameGeneratedSuccess struct {
	Name string
}

func (m AccountCharacterNameGeneratedSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterNameGeneratedSuccess
}

func (m AccountCharacterNameGeneratedSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Name), nil
}

func (m *AccountCharacterNameGeneratedSuccess) Deserialize(extra string) error {
	m.Name = extra

	return nil
}
