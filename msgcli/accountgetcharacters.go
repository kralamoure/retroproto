package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountGetCharacters struct{}

func (m AccountGetCharacters) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountGetCharacters
}

func (m AccountGetCharacters) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharacters) Deserialize(extra string) error {
	return nil
}
