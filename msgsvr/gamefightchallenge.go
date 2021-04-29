// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameFightChallenge struct{}

func (m GameFightChallenge) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameFightChallenge
}

func (m GameFightChallenge) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameFightChallenge) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
