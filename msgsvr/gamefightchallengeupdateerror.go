// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFightChallengeUpdateError struct{}

func NewGameFightChallengeUpdateError(extra string) (GameFightChallengeUpdateError, error) {
	var m GameFightChallengeUpdateError

	if err := m.Deserialize(extra); err != nil {
		return GameFightChallengeUpdateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFightChallengeUpdateError) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFightChallengeUpdateError
}

func (m GameFightChallengeUpdateError) MessageName() string {
	return "GameFightChallengeUpdateError"
}

func (m GameFightChallengeUpdateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightChallengeUpdateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
