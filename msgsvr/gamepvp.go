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
	return "", d1proto.ErrNotImplemented
}

func (m *GamePVP) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
