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
	return "", nil
}

func (m *AccountGetRandomCharacterName) Deserialize(extra string) error {
	return nil
}
