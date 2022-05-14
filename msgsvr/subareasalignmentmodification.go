// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SubareasAlignmentModification struct{}

func (m SubareasAlignmentModification) MessageId() retroproto.MsgSvrId {
	return retroproto.SubareasAlignmentModification
}

func (m SubareasAlignmentModification) MessageName() string {
	return "SubareasAlignmentModification"
}

func (m SubareasAlignmentModification) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubareasAlignmentModification) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
