// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameTurnReady struct{}

func (m GameTurnReady) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameTurnReady
}

func (m GameTurnReady) Serialized() (string, error) {
	return "", nil
}

func (m *GameTurnReady) Deserialize(extra string) error {
	return nil
}
