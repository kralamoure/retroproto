// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GamePositionStart struct{}

func (m GamePositionStart) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GamePositionStart
}

func (m GamePositionStart) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GamePositionStart) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
