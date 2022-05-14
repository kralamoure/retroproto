// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestConquestBonus struct{}

func NewConquestConquestBonus(extra string) (ConquestConquestBonus, error) {
	var m ConquestConquestBonus

	if err := m.Deserialize(extra); err != nil {
		return ConquestConquestBonus{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestConquestBonus) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestConquestBonus
}

func (m ConquestConquestBonus) MessageName() string {
	return "ConquestConquestBonus"
}

func (m ConquestConquestBonus) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestConquestBonus) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
