// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubareasList struct{}

func (m SubareasList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubareasList
}

func (m SubareasList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubareasList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
