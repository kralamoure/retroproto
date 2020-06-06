// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GamePVP struct{}

func (m GamePVP) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GamePVP
}

func (m GamePVP) Serialized() (string, error) {
	return "", nil
}

func (m *GamePVP) Deserialize(extra string) error {
	return nil
}
