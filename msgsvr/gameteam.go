// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTeam struct{}

func NewGameTeam(extra string) (GameTeam, error) {
	var m GameTeam

	if err := m.Deserialize(extra); err != nil {
		return GameTeam{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTeam) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTeam
}

func (m GameTeam) MessageName() string {
	return "GameTeam"
}

func (m GameTeam) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTeam) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
