// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameShowFightChallengeTarget struct{}

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
