package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountRequestRegionalVersion struct{}

func NewAccountRequestRegionalVersion(extra string) (AccountRequestRegionalVersion, error) {
	var m AccountRequestRegionalVersion

	if err := m.Deserialize(extra); err != nil {
		return AccountRequestRegionalVersion{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
