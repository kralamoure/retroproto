package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountRequestRegionalVersion struct{}

func (m AccountRequestRegionalVersion) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountRequestRegionalVersion
}

func (m AccountRequestRegionalVersion) Serialized() (string, error) {
	return "", nil
}

func (m *AccountRequestRegionalVersion) Deserialize(extra string) error {
	return nil
}
