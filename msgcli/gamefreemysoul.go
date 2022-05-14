// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFreeMySoul struct{}

func NewGameFreeMySoul(extra string) (GameFreeMySoul, error) {
	var m GameFreeMySoul

	if err := m.Deserialize(extra); err != nil {
		return GameFreeMySoul{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFreeMySoul) MessageId() retroproto.MsgCliId {
	return retroproto.GameFreeMySoul
}

func (m GameFreeMySoul) MessageName() string {
	return "GameFreeMySoul"
}

func (m GameFreeMySoul) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFreeMySoul) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
