// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameClearAllEffect struct{}

func NewGameClearAllEffect(extra string) (GameClearAllEffect, error) {
	var m GameClearAllEffect

	if err := m.Deserialize(extra); err != nil {
		return GameClearAllEffect{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameClearAllEffect) MessageId() retroproto.MsgSvrId {
	return retroproto.GameClearAllEffect
}

func (m GameClearAllEffect) MessageName() string {
	return "GameClearAllEffect"
}

func (m GameClearAllEffect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameClearAllEffect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
