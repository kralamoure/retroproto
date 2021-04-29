// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnStart struct{}

func (m GameTurnStart) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameTurnStart
}

func (m GameTurnStart) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnStart) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
