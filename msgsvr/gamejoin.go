// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameJoin struct{}

func NewGameJoin(extra string) (GameJoin, error) {
	var m GameJoin

	if err := m.Deserialize(extra); err != nil {
		return GameJoin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameJoin) MessageId() retroproto.MsgSvrId {
	return retroproto.GameJoin
}

func (m GameJoin) MessageName() string {
	return "GameJoin"
}

func (m GameJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
