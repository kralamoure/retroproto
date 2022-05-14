// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksQuickPing struct{}

func NewAksQuickPing(extra string) (AksQuickPing, error) {
	var m AksQuickPing

	if err := m.Deserialize(extra); err != nil {
		return AksQuickPing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksQuickPing) MessageId() retroproto.MsgCliId {
	return retroproto.AksQuickPing
}

func (m AksQuickPing) MessageName() string {
	return "AksQuickPing"
}

func (m AksQuickPing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksQuickPing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
