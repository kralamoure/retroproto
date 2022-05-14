// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SpellsSpellBoost struct{}

func NewSpellsSpellBoost(extra string) (SpellsSpellBoost, error) {
	var m SpellsSpellBoost

	if err := m.Deserialize(extra); err != nil {
		return SpellsSpellBoost{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsSpellBoost) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellBoost
}

func (m SpellsSpellBoost) MessageName() string {
	return "SpellsSpellBoost"
}

func (m SpellsSpellBoost) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SpellsSpellBoost) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
