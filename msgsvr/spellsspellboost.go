// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SpellsSpellBoost struct{}

func (m SpellsSpellBoost) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SpellsSpellBoost
}

func (m SpellsSpellBoost) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SpellsSpellBoost) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
