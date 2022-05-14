// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFightChallenge struct{}

func NewGameFightChallenge(extra string) (GameFightChallenge, error) {
	var m GameFightChallenge

	if err := m.Deserialize(extra); err != nil {
		return GameFightChallenge{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFightChallenge) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFightChallenge
}

func (m GameFightChallenge) MessageName() string {
	return "GameFightChallenge"
}

func (m GameFightChallenge) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightChallenge) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
