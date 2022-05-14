package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountGetCharactersForced struct{}

func (m AccountGetCharactersForced) MessageId() retroproto.MsgCliId {
	return retroproto.AccountGetCharactersForced
}

func (m AccountGetCharactersForced) MessageName() string {
	return "AccountGetCharactersForced"
}

func (m AccountGetCharactersForced) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharactersForced) Deserialize(extra string) error {
	return nil
}
