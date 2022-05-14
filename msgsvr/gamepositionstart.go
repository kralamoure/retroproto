// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GamePositionStart struct{}

func NewGamePositionStart(extra string) (GamePositionStart, error) {
	var m GamePositionStart

	if err := m.Deserialize(extra); err != nil {
		return GamePositionStart{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GamePositionStart) MessageId() retroproto.MsgSvrId {
	return retroproto.GamePositionStart
}

func (m GamePositionStart) MessageName() string {
	return "GamePositionStart"
}

func (m GamePositionStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePositionStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
