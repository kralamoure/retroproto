// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameShowFightChallengeTarget struct{}

func NewGameShowFightChallengeTarget(extra string) (GameShowFightChallengeTarget, error) {
	var m GameShowFightChallengeTarget

	if err := m.Deserialize(extra); err != nil {
		return GameShowFightChallengeTarget{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameShowFightChallengeTarget) MessageId() retroproto.MsgCliId {
	return retroproto.GameShowFightChallengeTarget
}

func (m GameShowFightChallengeTarget) MessageName() string {
	return "GameShowFightChallengeTarget"
}

func (m GameShowFightChallengeTarget) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameShowFightChallengeTarget) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
