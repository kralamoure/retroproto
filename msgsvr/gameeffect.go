// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameEffect struct{}

func (m GameEffect) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameEffect
}

func (m GameEffect) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameEffect) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
