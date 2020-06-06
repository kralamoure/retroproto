// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameEffect struct{}

func (m GameEffect) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameEffect
}

func (m GameEffect) Serialized() (string, error) {
	return "", nil
}

func (m *GameEffect) Deserialize(extra string) error {
	return nil
}
