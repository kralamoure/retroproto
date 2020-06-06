package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountGetCharactersForced struct{}

func (m AccountGetCharactersForced) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountGetCharactersForced
}

func (m AccountGetCharactersForced) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharactersForced) Deserialize(extra string) error {
	return nil
}
