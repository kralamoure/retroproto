// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountAttributeGiftToCharacter struct{}

func (m AccountAttributeGiftToCharacter) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountAttributeGiftToCharacter
}

func (m AccountAttributeGiftToCharacter) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountAttributeGiftToCharacter) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
