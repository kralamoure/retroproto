// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GameActionsFinish struct{}

func (m GameActionsFinish) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameActionsFinish
}

func (m GameActionsFinish) Serialized() (string, error) {
	return "", nil
}

func (m *GameActionsFinish) Deserialize(extra string) error {
	return nil
}
