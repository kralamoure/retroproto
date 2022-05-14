// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubareasList struct{}

func (m SubareasList) MessageId() retroproto.MsgSvrId {
	return retroproto.SubareasList
}

func (m SubareasList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubareasList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
