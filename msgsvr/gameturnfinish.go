// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameTurnFinish struct{}

func (m GameTurnFinish) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameTurnFinish
}

func (m GameTurnFinish) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameTurnFinish) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
