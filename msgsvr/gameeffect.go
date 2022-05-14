// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameEffect struct{}

func NewGameEffect(extra string) (GameEffect, error) {
	var m GameEffect

	if err := m.Deserialize(extra); err != nil {
		return GameEffect{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameEffect) MessageId() retroproto.MsgSvrId {
	return retroproto.GameEffect
}

func (m GameEffect) MessageName() string {
	return "GameEffect"
}

func (m GameEffect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEffect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
