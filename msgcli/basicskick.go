// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsKick struct{}

func NewBasicsKick(extra string) (BasicsKick, error) {
	var m BasicsKick

	if err := m.Deserialize(extra); err != nil {
		return BasicsKick{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsKick) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsKick
}

func (m BasicsKick) MessageName() string {
	return "BasicsKick"
}

func (m BasicsKick) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsKick) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
