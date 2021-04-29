// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameShowFightChallengeTarget struct{}

func (m GameShowFightChallengeTarget) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameShowFightChallengeTarget
}

func (m GameShowFightChallengeTarget) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameShowFightChallengeTarget) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
