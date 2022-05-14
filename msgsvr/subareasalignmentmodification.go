// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubareasAlignmentModification struct{}

func NewSubareasAlignmentModification(extra string) (SubareasAlignmentModification, error) {
	var m SubareasAlignmentModification

	if err := m.Deserialize(extra); err != nil {
		return SubareasAlignmentModification{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
