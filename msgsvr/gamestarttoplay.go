// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameStartToPlay struct{}

func NewGameStartToPlay(extra string) (GameStartToPlay, error) {
	var m GameStartToPlay

	if err := m.Deserialize(extra); err != nil {
		return GameStartToPlay{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameStartToPlay) MessageId() retroproto.MsgSvrId {
	return retroproto.GameStartToPlay
}

func (m GameStartToPlay) MessageName() string {
	return "GameStartToPlay"
}

func (m GameStartToPlay) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameStartToPlay) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
