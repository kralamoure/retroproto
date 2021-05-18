package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SpellsSpellForgetShow struct{}

func (m SpellsSpellForgetShow) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellForgetShow
}

func (m SpellsSpellForgetShow) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetShow) Deserialize(extra string) error {
	return nil
}
