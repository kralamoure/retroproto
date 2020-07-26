// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFightChallengeUpdateSuccess struct{}

func (m GameFightChallengeUpdateSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFightChallengeUpdateSuccess
}

func (m GameFightChallengeUpdateSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameFightChallengeUpdateSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
