// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnList struct{}

func (m GameTurnList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameTurnList
}

func (m GameTurnList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameTurnList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
