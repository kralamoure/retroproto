package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGetCharacters struct{}

func NewAccountGetCharacters(extra string) (AccountGetCharacters, error) {
	var m AccountGetCharacters

	if err := m.Deserialize(extra); err != nil {
		return AccountGetCharacters{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGetCharacters) MessageId() retroproto.MsgCliId {
	return retroproto.AccountGetCharacters
}

func (m AccountGetCharacters) MessageName() string {
	return "AccountGetCharacters"
}

func (m AccountGetCharacters) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetCharacters) Deserialize(extra string) error {
	return nil
}
