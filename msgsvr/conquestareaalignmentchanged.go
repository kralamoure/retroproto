// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestAreaAlignmentChanged struct{}

func NewConquestAreaAlignmentChanged(extra string) (ConquestAreaAlignmentChanged, error) {
	var m ConquestAreaAlignmentChanged

	if err := m.Deserialize(extra); err != nil {
		return ConquestAreaAlignmentChanged{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestAreaAlignmentChanged) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestAreaAlignmentChanged
}

func (m ConquestAreaAlignmentChanged) MessageName() string {
	return "ConquestAreaAlignmentChanged"
}

func (m ConquestAreaAlignmentChanged) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestAreaAlignmentChanged) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
