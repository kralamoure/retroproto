package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountGetServersList struct{}

func (m AccountGetServersList) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountGetServersList
}

func (m AccountGetServersList) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetServersList) Deserialize(extra string) error {
	return nil
}
