// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GameActionsFinish struct{}

func (m GameActionsFinish) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameActionsFinish
}

func (m GameActionsFinish) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameActionsFinish) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
