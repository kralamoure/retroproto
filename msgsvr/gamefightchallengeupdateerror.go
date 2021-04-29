// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFightChallengeUpdateError struct{}

func (m GameFightChallengeUpdateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFightChallengeUpdateError
}

func (m GameFightChallengeUpdateError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFightChallengeUpdateError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
