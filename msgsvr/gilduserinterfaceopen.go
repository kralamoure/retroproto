// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildUserInterfaceOpen struct{}

func NewGildUserInterfaceOpen(extra string) (GildUserInterfaceOpen, error) {
	var m GildUserInterfaceOpen

	if err := m.Deserialize(extra); err != nil {
		return GildUserInterfaceOpen{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildUserInterfaceOpen) MessageId() retroproto.MsgSvrId {
	return retroproto.GildUserInterfaceOpen
}

func (m GildUserInterfaceOpen) MessageName() string {
	return "GildUserInterfaceOpen"
}

func (m GildUserInterfaceOpen) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildUserInterfaceOpen) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
