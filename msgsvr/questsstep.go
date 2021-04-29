// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type QuestsStep struct{}

func (m QuestsStep) ProtocolId() d1proto.MsgSvrId {
	return d1proto.QuestsStep
}

func (m QuestsStep) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *QuestsStep) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
