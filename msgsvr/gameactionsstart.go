// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameActionsStart struct{}

func NewGameActionsStart(extra string) (GameActionsStart, error) {
	var m GameActionsStart

	if err := m.Deserialize(extra); err != nil {
		return GameActionsStart{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameActionsStart) MessageId() retroproto.MsgSvrId {
	return retroproto.GameActionsStart
}

func (m GameActionsStart) MessageName() string {
	return "GameActionsStart"
}

func (m GameActionsStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameActionsStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
