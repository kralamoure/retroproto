// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnReady struct{}

func (m GameTurnReady) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameTurnReady
}

func (m GameTurnReady) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnReady) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
