// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameChallenge struct{}

func (m GameChallenge) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameChallenge
}

func (m GameChallenge) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameChallenge) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
