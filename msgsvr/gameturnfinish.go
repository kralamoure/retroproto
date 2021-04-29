// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnFinish struct{}

func (m GameTurnFinish) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameTurnFinish
}

func (m GameTurnFinish) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameTurnFinish) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
