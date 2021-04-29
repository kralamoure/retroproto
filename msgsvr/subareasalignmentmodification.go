// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SubareasAlignmentModification struct{}

func (m SubareasAlignmentModification) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SubareasAlignmentModification
}

func (m SubareasAlignmentModification) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubareasAlignmentModification) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
