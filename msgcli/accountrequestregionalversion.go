package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountRequestRegionalVersion struct{}

func (m AccountRequestRegionalVersion) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountRequestRegionalVersion
}

func (m AccountRequestRegionalVersion) Serialized() (string, error) {
	return "", nil
}

func (m *AccountRequestRegionalVersion) Deserialize(extra string) error {
	return nil
}
