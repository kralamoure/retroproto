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
	return "", nil
}

func (m *GameTurnStart) Deserialize(extra string) error {
	return nil
}
