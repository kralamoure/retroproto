// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsSanctionMe struct{}

func NewBasicsSanctionMe(extra string) (BasicsSanctionMe, error) {
	var m BasicsSanctionMe

	if err := m.Deserialize(extra); err != nil {
		return BasicsSanctionMe{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsSanctionMe) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsSanctionMe
}

func (m BasicsSanctionMe) MessageName() string {
	return "BasicsSanctionMe"
}

func (m BasicsSanctionMe) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsSanctionMe) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
