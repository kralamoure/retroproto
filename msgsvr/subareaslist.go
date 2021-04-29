// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubareasList struct{}

func (m SubareasList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubareasList
}

func (m SubareasList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubareasList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
