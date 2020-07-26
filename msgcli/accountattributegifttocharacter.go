// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountAttributeGiftToCharacter struct{}

func (m AccountAttributeGiftToCharacter) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountAttributeGiftToCharacter
}

func (m AccountAttributeGiftToCharacter) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountAttributeGiftToCharacter) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
