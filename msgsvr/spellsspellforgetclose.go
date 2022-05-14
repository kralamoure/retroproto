package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SpellsSpellForgetClose struct{}

func (m SpellsSpellForgetClose) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellForgetClose
}

func (m SpellsSpellForgetClose) MessageName() string {
	return "SpellsSpellForgetClose"
}

func (m SpellsSpellForgetClose) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetClose) Deserialize(extra string) error {
	return nil
}
