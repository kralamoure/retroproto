// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type QuestsStep struct{}

func (m QuestsStep) MessageId() retroproto.MsgSvrId {
	return retroproto.QuestsStep
}

func (m QuestsStep) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestsStep) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
