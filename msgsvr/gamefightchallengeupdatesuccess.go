// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFightChallengeUpdateSuccess struct{}

func NewGameFightChallengeUpdateSuccess(extra string) (GameFightChallengeUpdateSuccess, error) {
	var m GameFightChallengeUpdateSuccess

	if err := m.Deserialize(extra); err != nil {
		return GameFightChallengeUpdateSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFightChallengeUpdateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFightChallengeUpdateSuccess
}

func (m GameFightChallengeUpdateSuccess) MessageName() string {
	return "GameFightChallengeUpdateSuccess"
}

func (m GameFightChallengeUpdateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightChallengeUpdateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
