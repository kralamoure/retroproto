package msgsvr

import (
	"fmt"

	"github.com/kralamoure/d1encoding"
)

type AccountCharacterNameGeneratedSuccess struct {
	Name string
}

func (m AccountCharacterNameGeneratedSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountCharacterNameGeneratedSuccess
}

func (m AccountCharacterNameGeneratedSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Name), nil
}

func (m *AccountCharacterNameGeneratedSuccess) Deserialize(extra string) error {
	m.Name = extra

	return nil
}
