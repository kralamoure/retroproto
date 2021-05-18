package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SpellsSpellForgetClose struct{}

func (m SpellsSpellForgetClose) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellForgetClose
}

func (m SpellsSpellForgetClose) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetClose) Deserialize(extra string) error {
	return nil
}
