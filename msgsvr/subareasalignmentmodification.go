// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubareasAlignmentModification struct{}

func (m SubareasAlignmentModification) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubareasAlignmentModification
}

func (m SubareasAlignmentModification) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubareasAlignmentModification) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
