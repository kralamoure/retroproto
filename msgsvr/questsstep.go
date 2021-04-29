// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type QuestsStep struct{}

func (m QuestsStep) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.QuestsStep
}

func (m QuestsStep) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *QuestsStep) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
