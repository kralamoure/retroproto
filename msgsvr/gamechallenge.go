// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameChallenge struct{}

func (m GameChallenge) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameChallenge
}

func (m GameChallenge) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameChallenge) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
