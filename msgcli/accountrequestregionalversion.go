package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountRequestRegionalVersion struct{}

func (m AccountRequestRegionalVersion) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountRequestRegionalVersion
}

func (m AccountRequestRegionalVersion) Serialized() (string, error) {
	return "", nil
}

func (m *AccountRequestRegionalVersion) Deserialize(extra string) error {
	return nil
}
