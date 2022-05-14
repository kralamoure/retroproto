// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameEnd struct{}

func NewGameEnd(extra string) (GameEnd, error) {
	var m GameEnd

	if err := m.Deserialize(extra); err != nil {
		return GameEnd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameEnd) MessageId() retroproto.MsgSvrId {
	return retroproto.GameEnd
}

func (m GameEnd) MessageName() string {
	return "GameEnd"
}

func (m GameEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
