// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameShowFightChallengeTarget struct{}

func (m GameShowFightChallengeTarget) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameShowFightChallengeTarget
}

func (m GameShowFightChallengeTarget) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameShowFightChallengeTarget) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
