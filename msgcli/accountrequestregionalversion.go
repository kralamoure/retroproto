package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountRequestRegionalVersion struct{}

func (m AccountRequestRegionalVersion) MessageId() retroproto.MsgCliId {
	return retroproto.AccountRequestRegionalVersion
}

func (m AccountRequestRegionalVersion) MessageName() string {
	return "AccountRequestRegionalVersion"
}

func (m AccountRequestRegionalVersion) Serialized() (string, error) {
	return "", nil
}

func (m *AccountRequestRegionalVersion) Deserialize(extra string) error {
	return nil
}
