// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountAttributeGiftToCharacter struct{}

func NewAccountAttributeGiftToCharacter(extra string) (AccountAttributeGiftToCharacter, error) {
	var m AccountAttributeGiftToCharacter

	if err := m.Deserialize(extra); err != nil {
		return AccountAttributeGiftToCharacter{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountAttributeGiftToCharacter) MessageId() retroproto.MsgCliId {
	return retroproto.AccountAttributeGiftToCharacter
}

func (m AccountAttributeGiftToCharacter) MessageName() string {
	return "AccountAttributeGiftToCharacter"
}

func (m AccountAttributeGiftToCharacter) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountAttributeGiftToCharacter) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
