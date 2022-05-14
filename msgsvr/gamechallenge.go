// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameChallenge struct{}

func NewGameChallenge(extra string) (GameChallenge, error) {
	var m GameChallenge

	if err := m.Deserialize(extra); err != nil {
		return GameChallenge{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameChallenge) MessageId() retroproto.MsgSvrId {
	return retroproto.GameChallenge
}

func (m GameChallenge) MessageName() string {
	return "GameChallenge"
}

func (m GameChallenge) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameChallenge) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
