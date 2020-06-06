// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsSpellBoost struct{}

func (m SpellsSpellBoost) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsSpellBoost
}

func (m SpellsSpellBoost) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellBoost) Deserialize(extra string) error {
	return nil
}
