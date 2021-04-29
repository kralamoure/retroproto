package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SpellsSpellForgetShow struct{}

func (m SpellsSpellForgetShow) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SpellsSpellForgetShow
}

func (m SpellsSpellForgetShow) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetShow) Deserialize(extra string) error {
	return nil
}
