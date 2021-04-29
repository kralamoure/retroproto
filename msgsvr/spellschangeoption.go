package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SpellsChangeOption struct {
	CanUseSeeAllSpell bool
}

func (m SpellsChangeOption) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SpellsChangeOption
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
