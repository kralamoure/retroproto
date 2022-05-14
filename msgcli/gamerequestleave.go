// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameRequestLeave struct{}

func NewGameRequestLeave(extra string) (GameRequestLeave, error) {
	var m GameRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return GameRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.GameRequestLeave
}

func (m GameRequestLeave) MessageName() string {
	return "GameRequestLeave"
}

func (m GameRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
