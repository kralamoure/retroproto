// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameChallenge struct{}

func (m GameChallenge) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameChallenge
}

func (m GameChallenge) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameChallenge) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
