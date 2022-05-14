// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameSetPlayerPosition struct{}

func NewGameSetPlayerPosition(extra string) (GameSetPlayerPosition, error) {
	var m GameSetPlayerPosition

	if err := m.Deserialize(extra); err != nil {
		return GameSetPlayerPosition{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameSetPlayerPosition) MessageId() retroproto.MsgCliId {
	return retroproto.GameSetPlayerPosition
}

func (m GameSetPlayerPosition) MessageName() string {
	return "GameSetPlayerPosition"
}

func (m GameSetPlayerPosition) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameSetPlayerPosition) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
