// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnOk struct{}

func NewGameTurnOk(extra string) (GameTurnOk, error) {
	var m GameTurnOk

	if err := m.Deserialize(extra); err != nil {
		return GameTurnOk{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnOk) MessageId() retroproto.MsgCliId {
	return retroproto.GameTurnOk
}

func (m GameTurnOk) MessageName() string {
	return "GameTurnOk"
}

func (m GameTurnOk) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnOk) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
