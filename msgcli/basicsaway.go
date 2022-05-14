// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAway struct{}

func NewBasicsAway(extra string) (BasicsAway, error) {
	var m BasicsAway

	if err := m.Deserialize(extra); err != nil {
		return BasicsAway{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAway) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAway
}

func (m BasicsAway) MessageName() string {
	return "BasicsAway"
}

func (m BasicsAway) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAway) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
