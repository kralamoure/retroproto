package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGetCharacters struct{}

func (m AccountGetCharacters) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountGetCharacters
}

func (m AccountGetCharacters) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharacters) Deserialize(extra string) error {
	return nil
}
