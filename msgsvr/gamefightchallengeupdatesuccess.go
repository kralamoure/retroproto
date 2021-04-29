// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFightChallengeUpdateSuccess struct{}

func (m GameFightChallengeUpdateSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFightChallengeUpdateSuccess
}

func (m GameFightChallengeUpdateSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFightChallengeUpdateSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
