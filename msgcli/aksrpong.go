// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksRPong struct{}

func NewAksRPong(extra string) (AksRPong, error) {
	var m AksRPong

	if err := m.Deserialize(extra); err != nil {
		return AksRPong{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksRPong) MessageId() retroproto.MsgCliId {
	return retroproto.AksRPong
}

func (m AksRPong) MessageName() string {
	return "AksRPong"
}

func (m AksRPong) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksRPong) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
