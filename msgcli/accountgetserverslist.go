package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGetServersList struct{}

func (m AccountGetServersList) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountGetServersList
}

func (m AccountGetServersList) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetServersList) Deserialize(extra string) error {
	return nil
}
