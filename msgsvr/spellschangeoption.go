package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SpellsChangeOption struct {
	CanUseSeeAllSpell bool
}

func (m SpellsChangeOption) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsChangeOption
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
