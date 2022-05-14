// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestGetAlignedBonus struct{}

func NewConquestGetAlignedBonus(extra string) (ConquestGetAlignedBonus, error) {
	var m ConquestGetAlignedBonus

	if err := m.Deserialize(extra); err != nil {
		return ConquestGetAlignedBonus{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestGetAlignedBonus) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestGetAlignedBonus
}

func (m ConquestGetAlignedBonus) MessageName() string {
	return "ConquestGetAlignedBonus"
}

func (m ConquestGetAlignedBonus) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestGetAlignedBonus) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
