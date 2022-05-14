// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameEnabledPVPMode struct{}

func NewGameEnabledPVPMode(extra string) (GameEnabledPVPMode, error) {
	var m GameEnabledPVPMode

	if err := m.Deserialize(extra); err != nil {
		return GameEnabledPVPMode{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameEnabledPVPMode) MessageId() retroproto.MsgCliId {
	return retroproto.GameEnabledPVPMode
}

func (m GameEnabledPVPMode) MessageName() string {
	return "GameEnabledPVPMode"
}

func (m GameEnabledPVPMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEnabledPVPMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
