// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnFinish struct{}

func NewGameTurnFinish(extra string) (GameTurnFinish, error) {
	var m GameTurnFinish

	if err := m.Deserialize(extra); err != nil {
		return GameTurnFinish{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnFinish) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnFinish
}

func (m GameTurnFinish) MessageName() string {
	return "GameTurnFinish"
}

func (m GameTurnFinish) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnFinish) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
