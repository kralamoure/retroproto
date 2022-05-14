// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameGameOver struct{}

func NewGameGameOver(extra string) (GameGameOver, error) {
	var m GameGameOver

	if err := m.Deserialize(extra); err != nil {
		return GameGameOver{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameGameOver) MessageId() retroproto.MsgSvrId {
	return retroproto.GameGameOver
}

func (m GameGameOver) MessageName() string {
	return "GameGameOver"
}

func (m GameGameOver) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameGameOver) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
