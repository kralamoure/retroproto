// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountSetNickname struct{}

func (m AccountSetNickname) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountSetNickname
}

func (m AccountSetNickname) Serialized() (string, error) {
	return "", nil
}

func (m *AccountSetNickname) Deserialize(extra string) error {
	return nil
}
