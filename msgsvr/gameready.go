// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameReady struct{}

func NewGameReady(extra string) (GameReady, error) {
	var m GameReady

	if err := m.Deserialize(extra); err != nil {
		return GameReady{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameReady) MessageId() retroproto.MsgSvrId {
	return retroproto.GameReady
}

func (m GameReady) MessageName() string {
	return "GameReady"
}

func (m GameReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
