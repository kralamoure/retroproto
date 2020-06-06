// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GamePositionStart struct{}

func (m GamePositionStart) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GamePositionStart
}

func (m GamePositionStart) Serialized() (string, error) {
	return "", nil
}

func (m *GamePositionStart) Deserialize(extra string) error {
	return nil
}
