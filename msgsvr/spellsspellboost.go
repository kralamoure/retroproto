// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SpellsSpellBoost struct{}

func (m SpellsSpellBoost) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellBoost
}

func (m SpellsSpellBoost) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SpellsSpellBoost) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
