package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGetCharactersForced struct{}

func NewAccountGetCharactersForced(extra string) (AccountGetCharactersForced, error) {
	var m AccountGetCharactersForced

	if err := m.Deserialize(extra); err != nil {
		return AccountGetCharactersForced{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
