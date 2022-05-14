// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGetRandomCharacterName struct{}

func NewAccountGetRandomCharacterName(extra string) (AccountGetRandomCharacterName, error) {
	var m AccountGetRandomCharacterName

	if err := m.Deserialize(extra); err != nil {
		return AccountGetRandomCharacterName{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGetRandomCharacterName) MessageId() retroproto.MsgCliId {
	return retroproto.AccountGetRandomCharacterName
}

func (m AccountGetRandomCharacterName) MessageName() string {
	return "AccountGetRandomCharacterName"
}

func (m AccountGetRandomCharacterName) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGetRandomCharacterName) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
