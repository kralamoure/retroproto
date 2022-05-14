// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountAttributeGiftToCharacter struct{}

func (m AccountAttributeGiftToCharacter) MessageId() retroproto.MsgCliId {
	return retroproto.AccountAttributeGiftToCharacter
}

func (m AccountAttributeGiftToCharacter) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountAttributeGiftToCharacter) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
