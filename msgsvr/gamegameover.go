// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameGameOver struct{}

func (m GameGameOver) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameGameOver
}

func (m GameGameOver) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameGameOver) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
