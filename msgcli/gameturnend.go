// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnEnd struct{}

func NewGameTurnEnd(extra string) (GameTurnEnd, error) {
	var m GameTurnEnd

	if err := m.Deserialize(extra); err != nil {
		return GameTurnEnd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnEnd) MessageId() retroproto.MsgCliId {
	return retroproto.GameTurnEnd
}

func (m GameTurnEnd) MessageName() string {
	return "GameTurnEnd"
}

func (m GameTurnEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
