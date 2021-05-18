// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameTurnStart struct{}

func (m GameTurnStart) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameTurnStart
}

func (m GameTurnStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
