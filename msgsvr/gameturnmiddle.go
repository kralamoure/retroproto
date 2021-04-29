// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnMiddle struct{}

func (m GameTurnMiddle) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameTurnMiddle
}

func (m GameTurnMiddle) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnMiddle) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
