// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameFightChallenge struct{}

func (m GameFightChallenge) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameFightChallenge
}

func (m GameFightChallenge) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameFightChallenge) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
