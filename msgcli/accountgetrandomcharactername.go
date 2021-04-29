// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGetRandomCharacterName struct{}

func (m AccountGetRandomCharacterName) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountGetRandomCharacterName
}

func (m AccountGetRandomCharacterName) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountGetRandomCharacterName) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
