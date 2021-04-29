// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameGameOver struct{}

func (m GameGameOver) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameGameOver
}

func (m GameGameOver) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameGameOver) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
