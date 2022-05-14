// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnStart struct{}

func NewGameTurnStart(extra string) (GameTurnStart, error) {
	var m GameTurnStart

	if err := m.Deserialize(extra); err != nil {
		return GameTurnStart{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnStart) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnStart
}

func (m GameTurnStart) MessageName() string {
	return "GameTurnStart"
}

func (m GameTurnStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
