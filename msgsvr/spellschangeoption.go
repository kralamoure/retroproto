package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsChangeOption struct {
	CanUseSeeAllSpell bool
}

func (m SpellsChangeOption) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsChangeOption
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
