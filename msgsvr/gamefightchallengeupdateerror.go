// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFightChallengeUpdateError struct{}

func (m GameFightChallengeUpdateError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameFightChallengeUpdateError
}

func (m GameFightChallengeUpdateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightChallengeUpdateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
