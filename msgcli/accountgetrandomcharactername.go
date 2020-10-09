// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountGetRandomCharacterName struct{}

func (m AccountGetRandomCharacterName) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountGetRandomCharacterName
}

func (m AccountGetRandomCharacterName) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountGetRandomCharacterName) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
