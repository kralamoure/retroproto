// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameActionsStart struct{}

func (m GameActionsStart) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameActionsStart
}

func (m GameActionsStart) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameActionsStart) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
