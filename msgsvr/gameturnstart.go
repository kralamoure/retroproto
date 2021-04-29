// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnStart struct{}

func (m GameTurnStart) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameTurnStart
}

func (m GameTurnStart) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameTurnStart) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
