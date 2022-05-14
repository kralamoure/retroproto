// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameFightChallengeUpdateSuccess struct{}

func (m GameFightChallengeUpdateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFightChallengeUpdateSuccess
}

func (m GameFightChallengeUpdateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightChallengeUpdateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
