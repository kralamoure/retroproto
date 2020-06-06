// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameClearAllEffect struct{}

func (m GameClearAllEffect) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameClearAllEffect
}

func (m GameClearAllEffect) Serialized() (string, error) {
	return "", nil
}

func (m *GameClearAllEffect) Deserialize(extra string) error {
	return nil
}
