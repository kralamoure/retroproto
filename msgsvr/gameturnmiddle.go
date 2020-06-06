// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnMiddle struct{}

func (m GameTurnMiddle) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameTurnMiddle
}

func (m GameTurnMiddle) Serialized() (string, error) {
	return "", nil
}

func (m *GameTurnMiddle) Deserialize(extra string) error {
	return nil
}
