// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsInvisible struct{}

func NewBasicsInvisible(extra string) (BasicsInvisible, error) {
	var m BasicsInvisible

	if err := m.Deserialize(extra); err != nil {
		return BasicsInvisible{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsInvisible) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsInvisible
}

func (m BasicsInvisible) MessageName() string {
	return "BasicsInvisible"
}

func (m BasicsInvisible) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsInvisible) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
