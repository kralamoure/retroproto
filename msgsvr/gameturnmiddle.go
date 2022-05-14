// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnMiddle struct{}

func NewGameTurnMiddle(extra string) (GameTurnMiddle, error) {
	var m GameTurnMiddle

	if err := m.Deserialize(extra); err != nil {
		return GameTurnMiddle{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnMiddle) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnMiddle
}

func (m GameTurnMiddle) MessageName() string {
	return "GameTurnMiddle"
}

func (m GameTurnMiddle) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnMiddle) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
