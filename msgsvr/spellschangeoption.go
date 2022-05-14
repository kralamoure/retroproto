package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SpellsChangeOption struct {
	CanUseSeeAllSpell bool
}

func NewSpellsChangeOption(extra string) (SpellsChangeOption, error) {
	var m SpellsChangeOption

	if err := m.Deserialize(extra); err != nil {
		return SpellsChangeOption{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsChangeOption) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsChangeOption
}

func (m SpellsChangeOption) MessageName() string {
	return "SpellsChangeOption"
}

func (m SpellsChangeOption) Serialized() (string, error) {
	canUseSeeAllSpell := "-"
	if m.CanUseSeeAllSpell {
		canUseSeeAllSpell = "+"
	}

	return canUseSeeAllSpell, nil
}

func (m *SpellsChangeOption) Deserialize(extra string) error {
	m.CanUseSeeAllSpell = extra == "+"

	return nil
}
