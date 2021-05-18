// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountGetRandomCharacterName struct{}

func (m AccountGetRandomCharacterName) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountGetRandomCharacterName
}

func (m AccountGetRandomCharacterName) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGetRandomCharacterName) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
