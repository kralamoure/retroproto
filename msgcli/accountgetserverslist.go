package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountGetServersList struct{}

func (m AccountGetServersList) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountGetServersList
}

func (m AccountGetServersList) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetServersList) Deserialize(extra string) error {
	return nil
}
