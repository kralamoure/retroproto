// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GameActionsFinish struct{}

func (m GameActionsFinish) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GameActionsFinish
}

func (m GameActionsFinish) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameActionsFinish) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
