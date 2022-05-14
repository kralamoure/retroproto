// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameAskDisablePVPMode struct{}

func NewGameAskDisablePVPMode(extra string) (GameAskDisablePVPMode, error) {
	var m GameAskDisablePVPMode

	if err := m.Deserialize(extra); err != nil {
		return GameAskDisablePVPMode{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameAskDisablePVPMode) MessageId() retroproto.MsgCliId {
	return retroproto.GameAskDisablePVPMode
}

func (m GameAskDisablePVPMode) MessageName() string {
	return "GameAskDisablePVPMode"
}

func (m GameAskDisablePVPMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameAskDisablePVPMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
