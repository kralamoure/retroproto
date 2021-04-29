package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGetCharactersForced struct{}

func (m AccountGetCharactersForced) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountGetCharactersForced
}

func (m AccountGetCharactersForced) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharactersForced) Deserialize(extra string) error {
	return nil
}
