// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameActionsStart struct{}

func (m GameActionsStart) MessageId() retroproto.MsgSvrId {
	return retroproto.GameActionsStart
}

func (m GameActionsStart) MessageName() string {
	return "GameActionsStart"
}

func (m GameActionsStart) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameActionsStart) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
