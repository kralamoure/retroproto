// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsWhoIs struct{}

func NewBasicsWhoIs(extra string) (BasicsWhoIs, error) {
	var m BasicsWhoIs

	if err := m.Deserialize(extra); err != nil {
		return BasicsWhoIs{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsWhoIs) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsWhoIs
}

func (m BasicsWhoIs) MessageName() string {
	return "BasicsWhoIs"
}

func (m BasicsWhoIs) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsWhoIs) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
