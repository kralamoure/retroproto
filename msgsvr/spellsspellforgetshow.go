// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsSpellForgetShow struct{}

func (m SpellsSpellForgetShow) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsSpellForgetShow
}

func (m SpellsSpellForgetShow) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetShow) Deserialize(extra string) error {
	return nil
}
