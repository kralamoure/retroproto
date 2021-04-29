package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SpellsSpellForgetClose struct{}

func (m SpellsSpellForgetClose) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SpellsSpellForgetClose
}

func (m SpellsSpellForgetClose) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetClose) Deserialize(extra string) error {
	return nil
}
