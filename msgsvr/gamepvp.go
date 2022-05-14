// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GamePVP struct{}

func NewGamePVP(extra string) (GamePVP, error) {
	var m GamePVP

	if err := m.Deserialize(extra); err != nil {
		return GamePVP{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GamePVP) MessageId() retroproto.MsgSvrId {
	return retroproto.GamePVP
}

func (m GamePVP) MessageName() string {
	return "GamePVP"
}

func (m GamePVP) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePVP) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
