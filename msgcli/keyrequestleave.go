// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type KeyRequestLeave struct{}

func NewKeyRequestLeave(extra string) (KeyRequestLeave, error) {
	var m KeyRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return KeyRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m KeyRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.KeyRequestLeave
}

func (m KeyRequestLeave) MessageName() string {
	return "KeyRequestLeave"
}

func (m KeyRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
