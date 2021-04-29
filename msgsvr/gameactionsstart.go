// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameActionsStart struct{}

func (m GameActionsStart) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameActionsStart
}

func (m GameActionsStart) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameActionsStart) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
