// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameLeave struct{}

func NewGameLeave(extra string) (GameLeave, error) {
	var m GameLeave

	if err := m.Deserialize(extra); err != nil {
		return GameLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.GameLeave
}

func (m GameLeave) MessageName() string {
	return "GameLeave"
}

func (m GameLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
