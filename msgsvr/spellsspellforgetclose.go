// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsSpellForgetClose struct{}

func (m SpellsSpellForgetClose) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsSpellForgetClose
}

func (m SpellsSpellForgetClose) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SpellsSpellForgetClose) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
