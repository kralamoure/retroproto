package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountGetCharacters struct{}

func (m AccountGetCharacters) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountGetCharacters
}

func (m AccountGetCharacters) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharacters) Deserialize(extra string) error {
	return nil
}
