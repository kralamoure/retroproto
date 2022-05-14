// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnReady struct{}

func NewGameTurnReady(extra string) (GameTurnReady, error) {
	var m GameTurnReady

	if err := m.Deserialize(extra); err != nil {
		return GameTurnReady{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnReady) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnReady
}

func (m GameTurnReady) MessageName() string {
	return "GameTurnReady"
}

func (m GameTurnReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
